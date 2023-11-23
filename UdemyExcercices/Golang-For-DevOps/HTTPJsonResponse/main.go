package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// google search: json to struct golang
// {"page":"words","input":"word1","words":["words"]}}
type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is invalid format: %s\n", err)
		os.Exit(1)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Status Code %d):\n %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	//Parse properties with struct

	var words Words
	err = json.Unmarshal(body, &words) //body: data []byte, &words v any pointer reference
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed\nPage: %s\nHTTP Status Code: %d\nWords: %v\n", words.Page, response.StatusCode, strings.Join(words.Words, ", "))
}
