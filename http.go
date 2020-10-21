package linkycore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CustomHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	HttpClient CustomHTTPClient
)

func GetJson(url string, target interface{}) error {
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

func PostJson(url string, body interface{}, target interface{}) error {
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
