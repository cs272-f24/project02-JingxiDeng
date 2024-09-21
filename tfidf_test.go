package main

import (
	"path"
	"testing"
)

func TestTfIdf(t *testing.T){
	tests := []struct{
		name string
		searchWord string
		seed string
		expected string
	}{
		{
			name: "Searching for Romeo",
			searchWord: "Romeo",
			seed: "top10/index.html",
			expected: "top10/The Project Gutenberg eBook of Romeo and Juliet, by William Shakespeare/index.html",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			server := MockServerHandler()
			defer server.Close()

			actual, err := TfIdf(test.searchWord, server.URL + path.Join("/", test.seed))
			if err != nil{
				t.Errorf("ERROR: Case %s\nTfIdf() returned: %v\n\n", test.name, err)
			}

			if actual != test.expected{
				t.Errorf("ERROR: Case %s\nExpected: %s\nActual: %s\n\n", test.name, test.expected, actual)
			}
		})
	}
}