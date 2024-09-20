package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestSearch(t *testing.T){
	tests := []struct{
		name string
		seed string
		searchWord string
		expected map[string]int
	}{
		{
			name: "Case: sceneI_30.0.html",
			seed: "rnj_files/sceneI_30.0.html",
			searchWord: "Verona",
			expected: map[string]int{
				"rnj_files/sceneI_30.0.html": 1,
			},
		},{
			name: "Case: sceneI_30.1.html",
			seed: "rnj_files/sceneI_30.1.html",
			searchWord: "Benvolio",
			expected: map[string]int{
				"rnj_files/sceneI_30.1.html": 26,
			},
		},{
			name: "Case: index.html",
			seed: "rnj_files/index.html",
			searchWord: "Romeo",
			expected: map[string]int{
				"rnj_files/index.html": 200,
				"rnj_files/sceneI_30.0.html":  2,
				"rnj_files/sceneI_30.1.html":  22,
				"rnj_files/sceneI_30.3.html":  2,
				"rnj_files/sceneI_30.4.html":  17,
				"rnj_files/sceneI_30.5.html":  15,
				"rnj_files/sceneII_30.2.html": 42,
				"rnj_files/sceneI_30.2.html":  15,
				"rnj_files/sceneII_30.0.html": 3,
				"rnj_files/sceneII_30.1.html": 10,
				"rnj_files/sceneII_30.3.html": 13,
			},
		},
		// {
		// 	name: "Case: top10",
		// 	seed: "top10/index.html",
		// 	searchWord: "Romeo",
		// 	expected: map[string]int{},
		// },
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			// Mock server serving the expected response
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Trim leading '/' from r.URL.Path
				filePath := strings.TrimPrefix(r.URL.Path, "/")
				
				file, err := os.Open(filePath)
				if err != nil {
					//t.Errorf("ERROR: %s is not a valid file path\n", filePath)
					return
				}
				defer file.Close()
		
				fileContent, err := io.ReadAll(file)
				if err != nil {
					t.Errorf("Error reading file: %v\n", err)
					return
				}
				w.Write(fileContent)
				
			})

			server := httptest.NewServer(handler)
			defer server.Close()

			// generate expected results:
			expectedResults := make(map[string]int)
			for key, val := range test.expected {
				expectedResults[key] = val
			}

			// check here
			// adding the mock server's url to the url provided in the test case
			actual, err := Search(server.URL + "/" + test.seed, test.searchWord)
			if err != nil {
				t.Errorf("ERROR: Search() returned \n%v\n", err)
			}
			
			if !reflect.DeepEqual(expectedResults, actual) {
				t.Errorf("\nERROR with %s\n Expected: %v\nActual: %v\n", test.name, expectedResults, actual)
			}
		})
	}
}