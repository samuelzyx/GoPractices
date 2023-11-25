package api

import (
	"encoding/json"
	"fmt"
	"io"
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

type WordsPage struct {
	Page
	Words
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

type Assignment1 struct {
	Words        []string           `json:"words"`
	Percentages  map[string]float32 `json:"percentages"`
	Special      []any              `json:"special"`
	ExtraSpecial []any              `json:"extraSpecial"`
}

func (a Assignment1) GetResponse() string {
	outPercentages := []string{}
	for word, percentage := range a.Percentages {
		outPercentages = append(outPercentages, fmt.Sprintf("%s (%v)", word, percentage))
	}
	outSpecial := []string{}
	for special := range a.Special {
		outSpecial = append(outSpecial, fmt.Sprintf("%v", special))
	}
	outExtraSpecial := []string{}
	for extraSpecial := range a.ExtraSpecial {
		outExtraSpecial = append(outExtraSpecial, fmt.Sprintf("%v", extraSpecial))
	}
	return fmt.Sprintf("Words: %s, Percentages: %s, Special: %s, ExtraSpecial: %s",
		strings.Join(a.Words, ", "), strings.Join(outPercentages, ", "), strings.Join(outSpecial, ", "), strings.Join(outExtraSpecial, ", "))
}

func (a API) DoGetRequest(requstURL string) (Response, error) {
	if _, err := url.ParseRequestURI(requstURL); err != nil {
		fmt.Printf("URL is invalid format: %s\n", err)
		os.Exit(1)
	}

	response, err := a.Client.Get(requstURL)

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
	case "assignment1":
		var assignment1 Assignment1
		err = json.Unmarshal(body, &assignment1) //body: data []byte, &words v any pointer reference
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Word unmarshal error: %s\n", err),
			}
		}

		return assignment1, nil

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
