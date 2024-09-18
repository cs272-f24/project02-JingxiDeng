package main

import (
	"github.com/kljensen/snowball"
)

func Search(seed, searchWord string) (map[string]int, error){
	searchResults, _,  err := Crawl(seed)
	if err != nil{
		return nil, err
	}

	stemmedSearchWord, err := snowball.Stem(searchWord, "english", true)
	if err != nil{
		return nil, err
	}
	frequencyMap := searchResults[stemmedSearchWord]
	
	return frequencyMap, nil
}
