package main

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

func main(){
	// I can serve a file as html content
	http.Handle("/", http.FileServer(http.Dir("static")))
	// I can serve a function as html server content
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request){
		searchTerm := r.URL.Query().Get("searchword")
		server := MockServerHandler()
		defer server.Close()

		// find the search result
		actual, err := TfIdf(searchTerm, server.URL + path.Join("/", "top10/index.html"))
		if err != nil{
			fmt.Println("ERROR with search")
		}
		// Un-decode the actual URL
		actual, err = url.PathUnescape(actual)
		if err != nil {
			fmt.Printf("ERROR: Failed to decode actual result: %v\n", err)
		}

		// display the search result
		w.Write([]byte("Most relevant document is: " + actual))
	})
	http.ListenAndServe(":8080", nil)
}