package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://example.com"

func main() {
	fmt.Println("Web Requests using golang")

	response, err := http.Get(url)

	if(err != nil){
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n\n\n",response)

	defer response.Body.Close()	// Imp to close the connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}