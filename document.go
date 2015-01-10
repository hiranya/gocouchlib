package gocouchlib

import (
	"net/http"
)

type Document struct {
	Id  string    `json:"_id"`
	Rev string    `json:"_rev"`
	Db  *Database `json:"-"`
}

func (doc *Document) endpoint(api string) string {
	return doc.Db.Server.FullUrl() + "/" + doc.Db.DbName + "/" + api
}

func (doc *Document) Exists() (bool, *CouchResponse) {

	headers := make(map[string][]string)

	if doc.Rev != "" {
		headers = map[string][]string{
			"If-None-Match": {"\"" + doc.Rev + "\""},
		}
	}

	couchResp, _ := httpClient.Head(doc.endpoint(doc.Id), headers)

	if (couchResp.StatusCode == http.StatusOK || couchResp.StatusCode == http.StatusNotModified) && couchResp.Headers.Get("ETag") != "" {
		doc.Rev = couchResp.Headers.Get("ETag")
	}

	return (couchResp.StatusCode == http.StatusOK || couchResp.StatusCode == http.StatusNotModified), couchResp
}
