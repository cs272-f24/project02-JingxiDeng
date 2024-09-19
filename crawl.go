package main

import (
	"fmt"
	"net/url"
	"github.com/kljensen/snowball"
)
/*
    This function updates the inverted index by inserting newly found words into the inverted index data structure
*/
func updateInvertedIndex(invertedIndex map[string]map[string]int, words []string, currentURL string) {
	// Add the extracted words into the inverted index
    for _, word := range words {
        // stem the word
        word, err := snowball.Stem(word, "english", true)
        if err != nil {
            fmt.Println(err)
            continue
        }
        urlMap, exists := invertedIndex[word]
        // If the word does not exist in the inverted index map, 
        // make a new hashmap and add it to the inverted index with the word as the new map's key
        if !exists {
            urlMap = make(map[string]int)
            invertedIndex[word] = urlMap
        }
        // Increment the frequency of the current word in the current URL by 1. 
        // The one-liner below works because if the current URL doesn't exist as a key, 
        // it automatically makes a new entry into the hashmap
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
	Crawl(): Given a seed URL, if the URL is a link and not a file path, then download the webpage,
		extract the words and URLs, and add all cleaned URLs found to a download queue and continue to crawl those URLs.
		If the seed is a file path, just access the .html data from the file through the os package
	@params: seed is the seed URL string that the method will crawl
	@returns: []string is the list of URLs that were crawled by the method
	@returns: error handles errors
*/

func Crawl(seed string) (map[string]map[string]int, []string, error) {
    invertedIndex := make(map[string]map[string]int) // words -> links -> frequency	

    queue := []string{seed}
    // make a hashset for the queue to speed up look-up times
    queueSet := make(map[string]struct{})
    queueSet[seed] = struct{}{}
    visited := make(map[string]struct{})

    var currentURL string
    for len(queue) > 0 {
        // Remove the current URL from the queue and parse it
        currentURL = queue[0]
        queue = queue[1:]
        delete(queueSet, currentURL)

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

        visited[currentURL] = struct{}{}
        extracted, err := Download(currentURL)
        if err != nil {
            return invertedIndex, nil, err
        }
        words, hrefs, err := Extract(extracted)
        if err != nil {
            return invertedIndex, nil, err
        }
        hrefs = Clean(currentURL, hrefs)

        updateInvertedIndex(invertedIndex, words, currentURL)

        addNewURLsToQueue(hrefs, currentURL, visited, &queue, &queueSet)
    }

    // Collect the visited URLs into a slice
    visitedURLs := make([]string, 0, len(visited))
    for url := range visited {
        visitedURLs = append(visitedURLs, url)
    }

    
    return invertedIndex, visitedURLs, nil
}


