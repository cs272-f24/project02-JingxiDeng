package main

/*
	Stop() will return stop words that are "meaningless", and stop words will not be added to the inverted index
	return true if a search word is a stop word
	return false if a search word is not a stop word
*/
func Stop(searchWord string, set map[string]struct{}) bool{
	_, exists := set[searchWord]
	return exists;
}