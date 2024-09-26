package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/kljensen/snowball"
)

type StopWords map[string]struct{}

/*
   Generate a hashset of stop words by reading from a JSON file
*/
func GenerateStopWords() (StopWords, error) {
	// Read the Stop words from "stopwords-en.json" and generate a stop word map
	fileContent, err := ioutil.ReadFile("stopwords-en.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var stopwords []string

	err = json.Unmarshal(fileContent, &stopwords)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil, err
	}

	// Make the set
	set := make(StopWords)
	for _, word := range stopwords {
		set[word] = struct{}{}
	}
	return set, nil
}

/*
    Removes the hostname and the prefix "/" from a string URL or a file path
*/
func removeHostname(fullURL string) (string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}

	// Get the path and trim the leading slash
	pathAndQuery := strings.TrimPrefix(parsedURL.EscapedPath(), "/")
	if parsedURL.RawQuery != "" {
		pathAndQuery += "?" + parsedURL.RawQuery
	}

	return pathAndQuery, nil
}

/*
   This function updates the inverted index by inserting newly found words into the inverted index data structure
*/
func updateInvertedIndex(idx *InvertedIndex, stopwords StopWords, words []string, currentURL string) {
	// Record the number of words inside of a particular document
	idx.docWordCount[currentURL] += len(words)

	// Add the extracted words into the inverted index
	for _, word := range words {
		// Check if the word is a stop word. If it is a stop word, skip to the next for loop run.
		if Stop(word, stopwords) {
			continue
		}
		// Stem the word
		stemmedWord, err := snowball.Stem(word, "english", true)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Retrieve the urlMap for the stemmed word
		urlMap, exists := idx.idx[stemmedWord]
		// If the word does not exist in the inverted index map, make a new hashmap and add it to the inverted index
		if !exists {
			urlMap = make(map[string]int)
			idx.idx[stemmedWord] = urlMap
		}
		// Increment the frequency of the current word in the current URL by 1
		urlMap[currentURL]++
	}
}

/*
    This function adds new, unvisited, and valid URL or file paths to the queue
*/
func addNewURLsToQueue(hrefs []string, currentURL string, visited map[string]struct{}, queue *[]string, queueSet *map[string]struct{}) {
	cleanedHrefs := Clean(currentURL, hrefs)
	for _, href := range cleanedHrefs {
		if _, seen := visited[href]; !seen && href != "INVALID HREF" {
			if _, inQueue := (*queueSet)[href]; !inQueue {
				*queue = append(*queue, href)
				(*queueSet)[href] = struct{}{}
			}
		}
	}
}

/*
	Crawl(): Given a seed URL, download the webpage, extract the words and URLs,
	add all cleaned URLs found to a download queue, and continue to crawl those URLs.

	@params: seed is the seed URL string that the method will crawl
	@returns: error for error handling
*/
func Crawl(idx *InvertedIndex, seed string) error {
	stopWords, err := GenerateStopWords()
	if err != nil {
		return err
	}

	queue := []string{seed}
	// Make a hashset for the queue to speed up look-up times
	queueSet := make(map[string]struct{})
	queueSet[seed] = struct{}{}
	visited := make(map[string]struct{})

	var currentURL string
	for len(queue) > 0 {
		// Remove the current URL from the queue and parse it
		currentURL = queue[0]
		queue = queue[1:]
		delete(queueSet, currentURL)

		if currentURL == "" {
			continue
		}

		parsedURL, err := url.Parse(currentURL)
		if err != nil {
			fmt.Printf("Skipping invalid URL: %s\n", currentURL)
			continue
		}
		currentURL = parsedURL.String()

		// If we've already visited the current URL, skip it.
		if _, alreadyVisited := visited[currentURL]; alreadyVisited || currentURL == "INVALID HREF" {
			continue
		}

		extracted, err := Download(currentURL)
		if err != nil {
			fmt.Printf("Error downloading URL %s: %v\n", currentURL, err)
			continue
		}
		words, hrefs, err := Extract(extracted)
		if err != nil {
			fmt.Printf("Error extracting data from URL %s: %v\n", currentURL, err)
			continue
		}
		hrefs = Clean(currentURL, hrefs)

		addNewURLsToQueue(hrefs, currentURL, visited, &queue, &queueSet)
		// Remove the hostname prefix for the current URL
		currentURL, err = removeHostname(currentURL)
		if err != nil {
			return err
		}
		visited[currentURL] = struct{}{}
		updateInvertedIndex(idx, stopWords, words, currentURL)
	}

	return nil
}
