package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const apiURL string = "https://slack.com/api/"

func main() {
	token := os.Getenv("SLACK_API_TOKEN")

	fmt.Println(token)

	response, err := http.Get(apiURL + "files.list?token=" + token)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)

}

func listFiles() {

}

func deleteFiles() {

}
