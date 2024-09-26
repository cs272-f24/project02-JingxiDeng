package main

import (
	"time"
)

type freq map[string]int

// InvertedIndex holds the index and document word counts
type InvertedIndex struct {
	idx          map[string]freq
	docWordCount map[string]int
}

func main() {
	idx := &InvertedIndex{
		idx:          make(map[string]freq),
		docWordCount: make(map[string]int),
	}

	go Crawl(idx, "http://localhost:8080/top10/index.html")
	go Serve(idx)
	
	for {
		time.Sleep(100)
	}
}
