package gocouchlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
)

// CouchResponse contains the elements of the HTTP response returned by the CouchDB Server
type CouchResponse struct {
	// Json payload returned by the CouchDB server as a response to the request
	Json JsonObj

	// HTTP Status code returned by the CouchDB server
	StatusCode int

	// HTTP Response Headers returned by the CouchDB server
	Headers http.Header
}

type JsonObj interface{}

var hc = &http.Client{}

type HttpClient struct {
}

var httpClient = &HttpClient{}

func (this *HttpClient) Get(url string, headers http.Header) (*CouchResponse, error) {
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

	return &CouchResponse{Json: getResponseJson(resp), StatusCode: resp.StatusCode, Headers: resp.Header}, nil
}

func (this *HttpClient) Head(url string, headers http.Header) (*CouchResponse, error) {
	fmt.Println("=> URL", url)

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		throwError(err)
	}

	if headers != nil {
		req.Header = headers
	}

	resp, err := hc.Do(req)
	if err != nil {
		throwError(err)
	}

	return &CouchResponse{Json: getResponseJson(resp), StatusCode: resp.StatusCode, Headers: resp.Header}, nil
}

func (this *HttpClient) Delete(url string) (*CouchResponse, error) {
	fmt.Println("URL: ", url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		throwError(err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := hc.Do(req)
	if err != nil {
		throwError(err)
	}

	return &CouchResponse{Json: getResponseJson(resp), StatusCode: resp.StatusCode, Headers: resp.Header}, nil
}

func (this *HttpClient) Put(url string, jsonObj JsonObj, headers http.Header) (*CouchResponse, error) {
	var req *http.Request
	var err error

	if jsonObj != nil {

		jsonBytes, err := json.Marshal(&jsonObj)
		if err != nil {
			throwError(err)
		}

		req, err = http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			throwError(err)
		}
	} else {
		req, err = http.NewRequest("PUT", url, nil)
		if err != nil {
			throwError(err)
		}
	}

	if headers != nil {
		req.Header = headers
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := hc.Do(req)
	if err != nil {
		throwError(err)
	}

	return &CouchResponse{Json: getResponseJson(resp), StatusCode: resp.StatusCode, Headers: resp.Header}, nil
}

func (this *HttpClient) Post(url string, jsonObj JsonObj) (*CouchResponse, error) {
	fmt.Println("=> utils.Save() entry: ")

	jsonBytes, err := json.Marshal(&jsonObj)
	if err != nil {
		throwError(err)
	}

	fmt.Println("=> utils.Save(): ", string(jsonBytes))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		throwError(err)
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := hc.Do(req)
	if err != nil {
		throwError(err)
	}

	return &CouchResponse{Json: getResponseJson(resp), StatusCode: resp.StatusCode, Headers: resp.Header}, nil
}

func throwError(err error) (JsonObj, error) {
	return nil, &CouchError{
		time.Now(), err.Error(),
	}
}

func getResponseJson(resp *http.Response) JsonObj {
	var jsonRespObj JsonObj

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		throwError(err)
	}

	//fmt.Println("=> [Log]: HTTP response dump:", string(body))

	err = json.Unmarshal(body, &jsonRespObj)
	if err != nil {
		throwError(err)
	}
	return jsonRespObj
}

func TrimEtag(etag string) string {
	return strings.Trim(etag, "\"")
}