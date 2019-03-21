package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// init the client
	client := &http.Client{}
	// init two variables to handle requests and errors from a NewRequest that "GET"'s the lateest listings using the api
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	// If there is a error print it to the console and exit
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "AUD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "64bc57e2-eb81-4dc3-86ee-f6d29ddd08f2")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

}
