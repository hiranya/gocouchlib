package gocouchlib

import (
	"net/http"
)

type Document struct {
	Id      string `json:"_id,omitempty"`
	Rev     string `json:"_rev,omitempty"`
	Deleted bool   `json:"_deleted,omitempty"`
	Json    JsonObj

	Db *Database `json:"-"`
}

func (doc *Document) dbEndpoint() string {
	return doc.Db.Server.FullUrl() + "/" + doc.Db.DbName
}

func (doc *Document) docEndpoint(params map[string]string) string {
	queryString := "?"
	if params != nil {
		for key, val := range params {
			queryString += key + "=" + val + "&"
		}
	}

	return doc.dbEndpoint() + "/" + doc.Id + queryString
}

func (doc *Document) Exists() (bool, *CouchResponse) {

	headers := make(http.Header)

	if doc.Rev != "" {
		headers.Add("If-None-Match", "\""+doc.Rev+"\"")
	}

	couchResp, _ := httpClient.Head(doc.docEndpoint(nil), headers)

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

	couchResp, _ := httpClient.Get(doc.docEndpoint(nil), headers)
	var json JsonObj = nil
	if couchResp.StatusCode == http.StatusOK || couchResp.StatusCode == http.StatusNotModified {
		json = couchResp.Json
		doc.Json = couchResp.Json
		doc.Rev = couchResp.Headers.Get("ETag")
	}

	return json, couchResp
}

func (doc *Document) Delete() (bool, *CouchResponse) {
	var params map[string]string
	params = make(map[string]string)

	params["rev"] = doc.Rev
	couchResp, _ := httpClient.Delete(doc.docEndpoint(params))

	if couchResp.StatusCode == http.StatusOK || couchResp.StatusCode == http.StatusAccepted {
		doc.Rev = couchResp.Headers.Get("ETag")
		doc.Deleted = true
	}

	return couchResp.StatusCode == http.StatusOK || couchResp.StatusCode == http.StatusAccepted, couchResp
}

func (doc *Document) Save() (bool, *CouchResponse) {
	var couchResp *CouchResponse

	if doc.Id == "" {
		couchResp, _ = httpClient.Post(doc.dbEndpoint(), doc.Json)
	} else {
		headers := make(http.Header)
		if doc.Rev != "" {
			headers.Add("If-Match", "\""+doc.Rev+"\"")
		}
		couchResp, _ = httpClient.Put(doc.docEndpoint(nil), doc.Json, headers)
	}

	if couchResp.StatusCode == http.StatusCreated || couchResp.StatusCode == http.StatusAccepted {
		doc.Rev = couchResp.Headers.Get("ETag")
	}

	return couchResp.StatusCode == http.StatusCreated || couchResp.StatusCode == http.StatusAccepted, couchResp
}
