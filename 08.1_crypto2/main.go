package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// init client
	client := &http.Client{}
	// Make a get request to the url that takes 3 parameters (method, ur, nul)
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	//if there is an error log it
	if err != nil {
		fmt.Println("Error requesting")
		os.Exit(1)
	}
	// Set header to accepts and the content-type
	req.Header.Set("Accepts", "application/json")
	// add header that specifiys the the api key, and the actual api key
	req.Header.Add("X-CMC_PRO_API_KEY", " 64bc57e2-eb81-4dc3-86ee-f6d29ddd08f2")
	// send request with the client
	resp, err := client.Do(req)
	// if error print then exit
	if err != nil {
		fmt.Println("Error getting response")
		os.Exit(1)
	}
	// print the response status
	fmt.Println(resp.Status)
	// create variable for the body ignoring the second input that reads the response body
	respBody, _ := ioutil.ReadAll(resp.Body)
	// print the var response body
	fmt.Println(string(respBody))
	//API KEY  64bc57e2-eb81-4dc3-86ee-f6d29ddd08f2
}
