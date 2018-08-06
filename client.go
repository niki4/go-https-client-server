package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {

	tr := &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     30 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Transport: tr,
		// TODO: Put here options for HTTP client headers, redirect, etc.
	}

	resp, err := client.PostForm(
		"http://localhost:9000/telemetry",
		url.Values{"key": {"value"}, "id": {"123"}})
	if err != nil {
		fmt.Println("Request error: ", err)
	}

	defer resp.Body.Close()

	fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error on parsing response body: ", err)
	}
	fmt.Println(body)
}
