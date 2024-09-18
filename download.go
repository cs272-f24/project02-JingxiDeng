package main

import (
	"io"
	"net/http"
)

func Download(url string) ([]byte, error){
	rsp, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	defer rsp.Body.Close()
	bts, err := io.ReadAll(rsp.Body)
	if err != nil{
		return nil, err
	}

	return bts, nil
}
