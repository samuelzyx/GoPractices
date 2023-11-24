package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response interface {
	GetResponse() string
}

type Page struct {
	Name string `json:"page"`
}

// google search: json to struct golang
// {"page":"words","input":"word1","words":["words"]}}
type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) GetResponse() string {
	return fmt.Sprintf("%s", strings.Join(w.Words, ", "))
}

type Ocurrence struct {
	Words map[string]int `json:"words"`
}

func (o Ocurrence) GetResponse() string {
	out := []string{}
	for word, ocurrence := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, ocurrence))
	}
	return fmt.Sprintf("%s", strings.Join(out, ", "))
}

func main() {
	var (
		requestURL string
		password   string
		parsedURL  *url.URL
		err        error
	)

	flag.StringVar(&requestURL, "url", "", "url to access")                       //Indicate argument -url to indicate in terminal as string
	flag.StringVar(&password, "password", "", "use a password to access our API") //Indicate argument -password to indicate in terminal as string

	flag.Parse()

	if parsedURL, err = url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("Validation error: URL is not valid %s\n", err)
		flag.Usage() //Use of flags in terminal in case wrong use of arguments
		os.Exit(1)
	}

	res, err := doRequest(parsedURL.String())
	if err != nil {
		if requestErr, ok := err.(RequestError); ok {
			fmt.Printf("Error: %s (HTTP Code: %d, Body: %s)\n", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
			os.Exit(1)
		}
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if res == nil {
		fmt.Printf("No response\n")
		os.Exit(1)
	}

	fmt.Printf("Response: %s\n", res.GetResponse())
}

func doRequest(requstURL string) (Response, error) {

	if _, err := url.ParseRequestURI(requstURL); err != nil {
		fmt.Printf("URL is invalid format: %s\n", err)
		os.Exit(1)
	}

	response, err := http.Get(requstURL)

	if err != nil {
		return nil, fmt.Errorf("Validation error: URL is not valid: %s\n", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid output (HTTP Status Code %d):\n %s\n", response.StatusCode, body)
	}

	if !json.Valid(body) {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("No valid JSON returned"),
		}
	}

	//Parse properties with struct

	var page Page

	err = json.Unmarshal(body, &page) //body: data []byte, &words v any pointer reference
	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("Page unmarshal error: %s\n", err),
		}
	}

	switch page.Name {
	case "words":
		var words Words
		err = json.Unmarshal(body, &words) //body: data []byte, &words v any pointer reference
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Word unmarshal error: %s\n", err),
			}
		}

		return words, nil

	case "occurrence":
		var ocurrence Ocurrence
		err = json.Unmarshal(body, &ocurrence) //body: data []byte, &words v any pointer reference
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Occurrence unmarshal error: %s\n", err),
			}
		}

		return ocurrence, nil

		/*
			if val, ok := ocurrence.Words["word1"]; ok {
				fmt.Printf("Found word1: %d\n", val)
			}

			for word, ocurrence := range ocurrence.Words {
				fmt.Printf("%s: %d\n", word, ocurrence)
			}

			fmt.Printf("JSON Parsed\nPage: %s\nHTTP Status Code: %d\nOcurrence: %v\n", page.Name, response.StatusCode, strings.Join(words.Words, ", "))
		*/
	default:
		fmt.Printf("Page not found\n")
	}

	return nil, nil

}
