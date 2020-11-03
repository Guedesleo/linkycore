package linkycore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type CustomHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	HttpClient CustomHTTPClient
)

// Get - Start a Get response and returns the http.Response without parse data
func Get(url string, contentType string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if contentType == "" {
		// any content type:
		req.Header.Set("Content-Type", "text/html; charset=utf-8")
	} else {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err2
	}
	bodyString := string(bodyBytes)
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAa:", bodyString)

	return resp, err
}

// GetJSON - Get and parse data in json format
func GetJSON(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// PostJSON - Send a Post request with JSON body and parse its result
func PostJSON(url string, body interface{}, target interface{}) error {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)

	fmt.Println("Result from post", bodyString)

	return json.NewDecoder(resp.Body).Decode(target)
}

// PostFormURLEncoded - Send a post request with form url encoded request body
func PostFormURLEncoded(url string, body url.Values, target interface{}) error {
	// data := url.Values{}
	// data.Set("name", "foo")
	// data.Set("surname", "bar")

	fmt.Println(body)
	fmt.Println(body.Encode())
	fmt.Println("encode ^^^")

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)

	fmt.Println("Result from PostFormUrlEncoded", bodyString)

	return json.Unmarshal(bodyBytes, target)
}
