package main

import (
	"math"
	"sort"
)

type Result struct{
	filepath string
	tfidfScore float64
}

/*
	Ranks the relavence of the results that come up from the search word
	TF (Term Frequency) = number of occurances of a word / number of words in a document
		- TF = frequencyMap[path] / docWordCount[path]
	IDF (Inverse Document Frequency) = log(number of documents / (number of documents that contain the search term + 1))
		- If a search word comes up in a small amount of documents, then those documents are more relavent.
		- IDF = log(number of docs / (docs containing the term + 1))
		-     = log(numOfDocs / (len(frequencyMap) + 1))

	TF IDF = TF * IDF

	@params searchWord is the word that the user typed into the search engine. TfIdf() will find the most relavent document for this searchWord.
	@params seed is the seed URL that the code will crawl to find search results
	@returns the file/URL path to the most relavent search result (the document with the highest TF-IDF score with the search term)
	@returns error for error handling
*/
func TfIdf(searchWord, seed string)(string, error){
	frequencyMap, docWordCount, err := Search(seed, searchWord)
	if(err != nil){
		return "", err
	}
	if len(frequencyMap) == 0{
		return "", nil
	}

	// calculate the IDF score
	idfScore := math.Log10(float64(len(docWordCount)) / float64((len(frequencyMap)+1)))

	var results []Result

	for key, val := range frequencyMap{
		tfScore := float64(val) / float64(docWordCount[key])

		// make a new Result struct that contains the file path and its TFIDF score
		results = append(results, Result{filepath: key, tfidfScore: (tfScore * idfScore)})
	}

	// sort the results in descending order by tfidfScore
	sort.Slice(results, func(i, j int)bool{
		if(results[i].tfidfScore != results[j].tfidfScore){
			return results[i].tfidfScore > results[j].tfidfScore
		}
		// if two Result(s) have the same tfidfScore, then compare by filepath name sorted in ascending order.
		// e.g: sceneII_30.0.html will be more relevant than sceneII_30.1.html
		return results[i].filepath < results[j].filepath
	})

	return results[0].filepath, nil
}