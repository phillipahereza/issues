package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ResponseItem struct {
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	Body      string `json:"body"`
	URL       string `json:"html_url"`
}

type Response struct {
	TotalCount        int             `json:"total_count"`
	IncompleteResults bool            `json:"incomplete_results"`
	Items             [] ResponseItem `json:"items"`
}

const (
	LANGUAGE = "go"
	LABEL    = "good-first-issue"
)

func main() {
	url := fmt.Sprintf("https://api.github.com/search/issues?q=label:%s+language:%s+state:open+comments:<2+created:>2019-07-01&sort=created&order=desc", LABEL, LANGUAGE)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response

	err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for i := 0; i < responseObject.TotalCount; i++ {
		item := responseObject.Items[i]
		fmt.Printf("%s [%s]\n", item.Title, item.URL)
	}
}
