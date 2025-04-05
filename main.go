package main

import (
	"fmt"
	"os"
	"net/url"
	"net/http"
	"log"
	"io"
	"encoding/json"
)

type Data struct {
	Id int
	Email string
	FirstName string
	LastName string
	Avatar string
}

type Words struct {
	Page int `json:"page"`
	Total int `json:"total"`
	Data []Data `json:"data"` 
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Introduce URL")
		os.Exit(1)
	}
	
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("wrong URL %s\n", err)
		os.Exit(1)
	}

	res, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	
	if err != nil {
		log.Fatal(err)
	}
	
	if res.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", res.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	err = json.Unmarshal(body,&words)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JSON Parsed:\n pages: %v\n words: %v\n", words.Page, words.Data) 
}
