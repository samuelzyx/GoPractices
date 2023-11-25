package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/samuelzyx/GoPractices/tree/main/UdemyExcercices/Golang-For-DevOps/HttpGetClient/pkg/api"
)

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

	apiInstance := api.New(api.Options{
		Password: password,
		LoginURL: parsedURL.Scheme + "://" + parsedURL.Host + "/login",
	})

	res, err := apiInstance.DoGetRequest(parsedURL.String())
	if err != nil {
		if requestErr, ok := err.(api.RequestError); ok {
			fmt.Printf("Error occurred: %s (HTTP Error: %d, Body: %s)\n", requestErr.Error(), requestErr.HTTPCode, requestErr.Body)
			os.Exit(1)
		}
		fmt.Printf("Error occurred: %s\n", err)
		os.Exit(1)
	}
	if res == nil {
		fmt.Printf("No response\n")
		os.Exit(1)
	}
	fmt.Printf("Response: %s\n", res.GetResponse())
}
