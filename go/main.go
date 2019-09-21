package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type response struct {
	Image string `json:"image"`
	Link  string `json:"link"`
}

type MyEvent struct {
	Floof string `json:"floof"`
}

// HandleRequest handle request.
func HandleRequest(event MyEvent) (string, error) {
	if event.Floof != "floof" {
		return "only accept floof as keyword", nil
	}

	url := fmt.Sprintf("https://randomfox.ca/%s", event.Floof)
	resp, errCall := http.Get(url)
	if errCall != nil {
		return "error from api", nil
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var v response
	err := json.Unmarshal(body, &v)
	if err != nil {
		return "format error from api", nil
	}

	return v.Image, nil
}

func main() {
	lambda.Start(HandleRequest)
}
