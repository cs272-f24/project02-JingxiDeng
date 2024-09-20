package main

/*
	Ranks the relavence of the results that come up from the search word
	TF (Term Frequency) = number of occurances of a word / number of words in a document
	IDF (Inverse Document Frequency) = number of documents in which search term occurs (Romeo appears in 9000 documents)
		- If a search word comes up in a small amount of documents, then those documents are more relavent.
		- IDF = log(number of docs / (docs containing the term + 1))

	TF IDF = TF * (1/IDF)

	use float64 (decimal number)

	TF := (float64)termCount / (float64)wordsInDoc
		- Collect all Term Frequencies and sort them to determine most relavent (largest) score
	
		make another map to record how many words are in a doc
		number of docs = that map's length
		map[url] = number of words in a doc
		len(map) = number of docs
*/
func TfIdf(){

}