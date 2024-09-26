package main

import (
	"github.com/kljensen/snowball"
)

/*
	Search() counts the occurrences of a specified search word across all documents,
	and also returns the total word count for each document.
	@params: idx is the inverted index.
	@params: searchWord is the word for which the method will count occurrences across the documents.
	@returns: A map that shows the frequency of the searchWord in each document.
	@returns: A map that shows the total word count for each document, which is useful for calculating
			  the Term Frequency-Inverse Document Frequency (TF-IDF) score.
	@returns: An error if any issues occur during processing.
*/
func Search(idx *InvertedIndex, searchWord string) (map[string]int, map[string]int, error) {
	stemmedSearchWord, err := snowball.Stem(searchWord, "english", true)
	if err != nil {
		return nil, nil, err
	}
	frequencyMap := idx.idx[stemmedSearchWord]
	if len(frequencyMap) == 0 {
		return nil, nil, nil
	}
	docWordCount := idx.docWordCount
	return frequencyMap, docWordCount, nil
}
