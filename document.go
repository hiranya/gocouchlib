package gocouchlib

import (
	"net/http"
)

type Document struct {
	Id   string `json:"_id"`
	Rev  string `json:"_rev"`
	Json JsonObj

	Db *Database `json:"-"`
}

func (doc *Document) endpoint(api string) string {
	return doc.Db.Server.FullUrl() + "/" + doc.Db.DbName + "/" + api
}

func (doc *Document) Exists() (bool, *CouchResponse) {

	headers := make(http.Header)

	if doc.Rev != "" {
		headers.Add("If-None-Match", "\""+doc.Rev+"\"")
	}

	couchResp, _ := httpClient.Head(doc.endpoint(doc.Id), headers)

	// Set doc.Rev using the ETag on the response if it is currently empty
	if (couchResp.StatusCode == http.StatusOK || couchResp.StatusCode == http.StatusNotModified) && doc.Rev == "" && couchResp.Headers.Get("ETag") != "" {
		doc.Rev = couchResp.Headers.Get("ETag")
	}

	return (couchResp.StatusCode == http.StatusOK || couchResp.StatusCode == http.StatusNotModified), couchResp
}

func (doc *Document) Get() (JsonObj, *CouchResponse) { 

	headers := make(http.Header)

	if doc.Rev != "" {
		headers.Add("If-None-Match", "\""+doc.Rev+"\"")
	}

	couchResp, _ := httpClient.Get(doc.endpoint(doc.Id), headers)
	var json JsonObj = nil
	if couchResp.StatusCode == http.StatusOK || couchResp.StatusCode == http.StatusNotModified {
		json = couchResp.Json
		doc.Json = couchResp.Json
	}

	return json, couchResp
}
