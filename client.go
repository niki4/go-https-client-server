package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	tr := &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     30 * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		// TODO: Put here options for HTTP client headers, redirect, etc.
	}

	resp, err := client.PostForm(
		"https://localhost:9001/telemetry",
		url.Values{"key": {"value"}, "id": {"123"}})
	if err != nil {
		log.Fatalf("Request error: %s", err)
	}

	defer resp.Body.Close()

	fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error on parsing response body: ", err)
	}
	fmt.Println("Got message from server:", string(body))
}
