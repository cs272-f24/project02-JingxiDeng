package main

import (
	"log"
	"reflect"
	"testing"

	
)

func TestExtract(t *testing.T){
	tests := []struct{
		name string
		doc []byte
		words []string
		hrefs []string
	}{
		{
			name: "General Case",
			doc: []byte(`<!DOCTYPE html><html><body>
					<a href = "/something">sth<a>
					<a href = "/another">another<a>
					<a href = "https://www.example.com/">example<a>
				</body>
			</html>`), 
			words: []string{"sth", "another", "example"},
			hrefs: []string{"/something", "/another", "https://www.example.com/"},
		},{
			name: "Lab02 Example Case",
			doc: []byte(`<!DOCTYPE html>
			<html>
				<head>
					<title>CS272 | Welcome</title>
				</head>
				<body>
					<p>Hello World!</p>
					<p>Welcome to <a href="https://cs272-f24.github.io/">CS272</a>!</p>
				</body>
			</html>`),
			words: []string{"CS272", "Welcome", "Hello", "World", "Welcome", "to", "CS272"},
			hrefs: []string{"https://cs272-f24.github.io/"},
		},{
			name: "Blank Case",
			doc: []byte(``),
			words: nil,
			hrefs: nil,
		},
	}	

	for _, test := range tests{
		actualWords, actualHrefs, err := Extract(test.doc)
		if err != nil{
			log.Fatalf("ERROR. Extract(%v) returned %v, %v", test.doc, actualWords, actualHrefs)
		}

		if !reflect.DeepEqual(actualWords, test.words){
			t.Errorf("Failed Test Case: %s\nExpected: %q\nActual: %q\n\n", test.name, test.words, actualWords)
		}
		if !reflect.DeepEqual(actualHrefs, test.hrefs){
			t.Errorf("Failed Test Case: %s\nExpected: %q\nActual: %q\n\n", test.name, test.hrefs, actualHrefs)
		}
	}
}