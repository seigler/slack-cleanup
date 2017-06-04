package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Files contains the list of files we want to clean
type Files struct {
	FileList []File `json:"files"`
}

// File is an individual file we wish to remove from Slack
type File struct {
	ID string `json:"id"`
}

func main() {

	req, err := http.NewRequest("GET", "https://slack.com/api/files.list", nil)
	if err != nil {
		log.Fatal(err)
	}

	tsNow := time.Now()
	tsTo := tsNow.AddDate(0, -14, 0)
	tsUnix := tsTo.Unix()

	q := req.URL.Query()
	q.Add("token", os.Getenv("SLACK_API_TOKEN"))
	q.Add("ts_to", fmt.Sprint(tsUnix))
	q.Add("count", "1000")
	req.URL.RawQuery = q.Encode()

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var files Files
	if err := json.Unmarshal(content, &files); err != nil {
		log.Fatal(err)
	}

	for _, file := range files.FileList {
		req, err := http.NewRequest("GET", "https://slack.com/api/files.delete", nil)
		if err != nil {
			log.Fatal(err)
		}

		q := req.URL.Query()
		q.Add("token", os.Getenv("SLACK_API_TOKEN"))
		q.Add("file", file.ID)
		req.URL.RawQuery = q.Encode()

		httpClient := &http.Client{}
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("File", file.ID, "successfully deleted")
	}

}
