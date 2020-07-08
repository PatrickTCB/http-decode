package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		encodedValue := os.Args[1]
		decodedValue, err := url.QueryUnescape(encodedValue)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
		fmt.Println(decodedValue)
	} else {
		fmt.Println("No text to decode submitted.")
	}
}
