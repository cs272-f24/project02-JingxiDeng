package main

import (
	"net/url"
	"path"
	"testing"
)

func TestTfIdf(t *testing.T){
	tests := []struct{
		name string
		searchWord string
		expected string
	}{
		{
			name: "Searching for rabbit",
			searchWord: "rabbit",
			expected: "top10/The Project Gutenberg eBook of Alice’s Adventures in Wonderland, by Lewis Carroll/chap04.html",
		},{
			name: "Searching for prince",
			searchWord: "prince",
			expected: "top10/The Project Gutenberg eBook of The Prince, by Nicolo Machiavelli/chap00.html",
		},{
			name: "Searching for romeo",
			searchWord: "romeo",
			expected: "top10/The Project Gutenberg eBook of Romeo and Juliet, by William Shakespeare/sceneII_30.0.html",
		},
	}


	server := MockServerHandler()
	defer server.Close()
			
	// Initialize the inverted index
	idx := &InvertedIndex{
		idx:          make(map[string]freq),
		docWordCount: make(map[string]int),
	}
	Crawl(idx, server.URL + path.Join("/", "top10/index.html"))

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			actual, err := TfIdf(idx, test.searchWord)
			if err != nil{
				t.Errorf("%v\n", err)
			}

			// Un-decode the actual URL
			actual, err = url.PathUnescape(actual)
			if err != nil {
				t.Fatalf("ERROR: Failed to decode actual result: %v\n", err)
			}

			if actual != test.expected{
				t.Errorf("ERROR: Case %s\nExpected: %s\nActual: %s\n\n", test.name, test.expected, actual)
			}
		})
	}
}