package main

import (
	"math"
	"sort"
)

type Result struct {
	filepath   string
	tfidfScore float64
}

/*
	Ranks the relevance of the results that come up from the search word
	TF (Term Frequency) = number of occurrences of a word / number of words in a document
	IDF (Inverse Document Frequency) = log(number of documents / (number of documents that contain the search term + 1))

	@params idx is the inverted index.
	@params searchWord is the word that the user typed into the search engine.
	TfIdf() will find the most relevant document for this searchWord.
	@returns the file/URL path to the most relevant search result (the document with the highest TF-IDF score with the search term)
	@returns error for error handling
*/
func TfIdf(idx *InvertedIndex, searchWord string) (string, error) {
	frequencyMap, docWordCount, err := Search(idx, searchWord)
	if err != nil {
		return "", err
	}
	if len(frequencyMap) == 0 {
		return "", nil
	}

	// Calculate the IDF score
	idfScore := math.Log10(float64(len(docWordCount)) / float64((len(frequencyMap) + 1)))

	var results []Result

	for key, val := range frequencyMap {
		tfScore := float64(val) / float64(docWordCount[key])

		// Make a new Result struct that contains the file path and its TF-IDF score
		results = append(results, Result{filepath: key, tfidfScore: (tfScore * idfScore)})
	}

	// Sort the results in descending order by tfidfScore
	sort.Slice(results, func(i, j int) bool {
		if results[i].tfidfScore != results[j].tfidfScore {
			return results[i].tfidfScore > results[j].tfidfScore
		}
		// If two Results have the same tfidfScore, then compare by filepath name sorted in ascending order
		return results[i].filepath < results[j].filepath
	})

	return results[0].filepath, nil
}
