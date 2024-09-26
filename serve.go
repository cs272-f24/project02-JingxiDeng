package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func Serve(idx *InvertedIndex) {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("static")))

	http.Handle("/top10/", http.StripPrefix("/top10/", http.FileServer(http.Dir("./top10"))))

	// Handle search requests
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		searchTerm := r.URL.Query().Get("searchword")
		//server := MockServerHandler()
		//defer server.Close()

		// Find the search result
		actual, err := TfIdf(idx, searchTerm)
		if err != nil {
			fmt.Println("ERROR with search:", err)
		}
		// Un-decode the actual URL
		actual, err = url.PathUnescape(actual)
		if err != nil {
			fmt.Printf("ERROR: Failed to decode actual result: %v\n", err)
		}

		// Display the search result
		w.Write([]byte("Most relevant document is: " + actual))
		//fmt.Println(server.URL)
	})

	http.ListenAndServe(":8080", nil)
}
