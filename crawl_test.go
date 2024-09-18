package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	//"reflect"
	"slices"
	"strings"
	"testing"
)

/*
	arraysAreEqual compares if two arrays are equal in content, ignoring the order in which the content is sorted in.
*/
func arraysAreEqual(expected, actual []string) bool {
    if len(expected) != len(actual) {
        return false
    }
    set := make(map[string]struct{}, len(expected))
    for _, item := range expected {
        set[item] = struct{}{}
    }
    for _, item := range actual {
        if _, exists := set[item]; !exists {
            return false
        }
    }
    return true
}


func TestCrawl(t *testing.T){
	tests := []struct{
		name string
		seed string
		mockData []byte
		expected []string // 'expected' will be the list of URLs that Crawl() visited, meaning it's the equivalence of Crawl()'s 'visited' array
	}{
		{
			name: "Case: /index.html",
			seed: "rnj_files/index.html",
			expected: []string{
				"rnj_files/index.html",
				"rnj_files/sceneI_30.0.html",
				"rnj_files/sceneI_30.1.html",
				"rnj_files/sceneI_30.2.html",
				"rnj_files/sceneI_30.3.html",
				"rnj_files/sceneI_30.4.html",
				"rnj_files/sceneI_30.5.html",
				"rnj_files/sceneII_30.0.html",
				"rnj_files/sceneII_30.1.html",
				"rnj_files/sceneII_30.2.html",
				"rnj_files/sceneII_30.3.html",
			},
		},
	}

	// make mock server and run tests
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			// Mock server serving the expected response
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Trim leading '/' from r.URL.Path
				filePath := strings.TrimPrefix(r.URL.Path, "/")
				
				// Check if the file is in the expected list
				if slices.Contains(test.expected, filePath) {
					file, err := os.Open(filePath)
					if err != nil {
						t.Errorf("ERROR: %s is not a valid file path\n", filePath)
						return
					}
					defer file.Close()
			
					fileContent, err := io.ReadAll(file)
					if err != nil {
						t.Errorf("Error reading file: %v\n", err)
						return
					}
					w.Write(fileContent)
				} else {
					// Handle requests for files not in the expected list
					http.NotFound(w, r)
				}
			})

			server := httptest.NewServer(handler)
			defer server.Close()

			// generate expected results:
			expectedURLs := make([]string, len(test.expected))
			for i, p := range test.expected {
				expectedURLs[i] = server.URL + "/" + p
			}

			// adding the mock server's url to the url provided in the test case
			_, actual, err := Crawl(server.URL + "/" + test.seed)
			if err != nil {
				t.Errorf("ERROR: Crawl() returned %v\n", err)
			}
			
			if !arraysAreEqual(expectedURLs, actual) {
				t.Errorf("\nERROR with %s\n Expected: %v\nActual: %v\n", test.name, expectedURLs, actual)
			}
		})
	}
}