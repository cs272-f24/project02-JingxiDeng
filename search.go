package main

import (
	"github.com/kljensen/snowball"
)

/*
	Search() counts the occurrences of a specified search word across all documents, 
	and also returns the total word count for each document.
	@params: seed is the file path or URL that the Crawl() method will use to start crawling.
	@params: searchWord is the word for which the method will count occurrences across the documents.
	@returns: A map that shows the frequency of the searchWord in each document.
	@returns: A map that shows the total word count for each document, which is useful for calculating 
			  the Term Frequency-Inverse Document Frequency (TF-IDF) score.
	@returns: An error if any issues occur during processing.
*/
func Search(seed, searchWord string) (map[string]int, map[string]int, error) {
	searchResults, docWordCount, _, err := Crawl(seed)
	if err != nil {
		return nil, nil, err
	}

	stemmedSearchWord, err := snowball.Stem(searchWord, "english", true)
	if err != nil {
		return nil, nil, err
	}
	frequencyMap := searchResults[stemmedSearchWord]

	return frequencyMap, docWordCount, nil
}
