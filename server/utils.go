package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// HTTPGetRequest make simple HTTP request with method GET
func HTTPGetRequest(reqURL string) ([]byte, error, int) {
	client := &http.Client{Timeout: 30 * time.Second}

	resp, err := client.Get(reqURL)
	if err != nil {
		return nil, err, resp.StatusCode
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err, resp.StatusCode
	}

	return body, nil, resp.StatusCode
}

func checkError(err error) {
	if err != nil {
		log.Printf("fatal error: %s", err)
	}
}
