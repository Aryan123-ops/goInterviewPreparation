package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		log.Fatal(err.Error())

	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(responseData))
}
