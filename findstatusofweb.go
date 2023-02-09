package main

import (
	"fmt"
	"net/http"
)

// program to print the status of any web address ! Is it live or not ?
// for array of websites
func main() {
	url := []string{"https://go.dev/play/", "https://www.facebook.com/"}
	for _, v := range url {
		response, err := http.Get(v)
		if err != nil {
			fmt.Println("error")
		} else {
			fmt.Println(response.StatusCode)
		}
	}
}

// for a single website
// func main() {
// 	url := "https://go.dev/play/"
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	} else {
// 		fmt.Println(response.StatusCode)
// 	}
// }
