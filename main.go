package main

import (
	"fmt"
	"os"
)


func directoryExists(path string) bool {
    info, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return info.IsDir()
}

// wget -r -np -nH --cut-dirs=1 https://cs272-f24.github.io/top10         
func main(){
	// searchResults, err := Search(os.Args[1], os.Args[2])
	// if(err != nil){
	// 	fmt.Println(err)
	// }
	// fmt.Println(searchResults)

	_, err := os.Open(os.Args[1])
	if err != nil{
		fmt.Println("Not found")
	}else{
		fmt.Println("Found")
	}
}