package main

//import "math"

type DocFrequency struct{
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
*/
func TfIdf(searchWord, seed string)(string, error){
	// frequencyMap, docWordCount, err := Search(seed, searchWord)
	// if(err != nil){
	// 	return "", err
	// }

	// idfScore := math.Log(float64(len(docWordCount)) / float64((len(frequencyMap)+1)))



	return "", nil
}