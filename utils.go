package gocouchlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CouchResponse struct {
	// json payload
	Json JsonObj

	// error responses
	StatusCode int
//	Error      string
//	Reason     string
}

type JsonObj interface{}

var hc = &http.Client{}

type HttpClient struct{}

func (this *HttpClient) Get(url string) (*CouchResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		throwError(err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := hc.Do(req)
	if err != nil {
		throwError(err)
	}
	defer resp.Body.Close()

	var jsonObj JsonObj

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		throwError(err)
	}

	//fmt.Println("=> [Log]: HTTP response dump:", string(body))

	err = json.Unmarshal(body, &jsonObj)
	if err != nil {
		throwError(err)
	}

	return &CouchResponse{Json: jsonObj, StatusCode: resp.StatusCode}, nil
}

func (this *HttpClient) Head(url string) (*CouchResponse, error) {
	fmt.Println("=> URL", url)

	resp, err := hc.Head(url)
	if err != nil {
		throwError(err)
	}

	var jsonObj JsonObj

	return &CouchResponse{Json: jsonObj, StatusCode: resp.StatusCode}, nil
}

func (this *HttpClient) Delete(url string) (*CouchResponse, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		throwError(err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := hc.Do(req)
	if err != nil {
		throwError(err)
	}

	return &CouchResponse{Json: nil, StatusCode: resp.StatusCode}, nil
}

func (this *HttpClient) Put(url string) (*CouchResponse, error) {
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		throwError(err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := hc.Do(req)
	if err != nil {
		throwError(err)
	}
	
	return &CouchResponse{Json: nil, StatusCode: resp.StatusCode}, nil
}

func throwError(err error) (JsonObj, error) {
	return nil, &CouchError{
		time.Now(), err.Error(),
	}
}
