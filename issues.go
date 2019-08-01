package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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


func getMinDate(days int) string {
	today := time.Now()
	minDate := today.AddDate(0, 0, -1 * days)
	return fmt.Sprintf(minDate.Format("2006-01-02"))
}

func FetchIssues(language, label string, days int) (Response, error){
	minDate := getMinDate(days)
	url := fmt.Sprintf("https://api.github.com/search/issues?q=label:%s+language:%s+state:open+comments:<2+created:>%s&sort=created&order=desc", label, language, minDate)

	response, err := http.Get(url)

	if err != nil {
		return Response{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	var responseObject Response

	err = json.Unmarshal(responseData, &responseObject)

	if err != nil {
		return Response{}, err
	}

	return responseObject, nil

}

func PrintResponse(writer io.Writer, responseObject Response) {
	for i := 0; i < responseObject.TotalCount; i++  {
		item := responseObject.Items[i]
		fmt.Fprintf(writer, "%s\n", item.URL)
	}
}

func main() {
	var language string
	var label string
	var days int

	app := cli.NewApp()
	app.Name = "Issue Finder"
	app.Usage = "Need to contribute to OpenSource? Find fresh issues"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name:        "language, l",
			Value:       "",
			Usage:       "search for issues in this language",
			Destination: &language,
		},
		cli.StringFlag{
			Name:        "label, b",
			Value:       "good-first-issue",
			Usage:       "Search for issues with this label",
			Destination: &label,
		},
		cli.IntFlag{
			Name:        "days, d",
			Value:       30,
			Usage:       "Search for issues created within the last n days",
			Destination: &days,
		},

	}

	app.Action = func(c *cli.Context) error {
		if language == "" {
			_ = cli.ShowAppHelp(c)
			os.Exit(0)
		}

		responseObject, err := FetchIssues(language, label, days)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		PrintResponse(os.Stdout, responseObject)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
