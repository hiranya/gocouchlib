package gocouchlib

import (
	"net/http"
)

type Document struct {
	Id  string
	Rev string
	Db  *Database
}

func (doc *Document) endpoint(api string) string {
	return doc.Db.Server.FullUrl() + "/" + doc.Db.DbName + "/" + api
}

func (doc *Document) Exists() (bool, *CouchResponse) {
	couchResp, _ := httpClient.Head(doc.endpoint(doc.Id))
	return couchResp.StatusCode == http.StatusOK, couchResp
}

//func (doc *Document) Exists() (bool, *CouchResponse) {
//	couchResp, _ := httpClient.Head(doc.endpoint(doc.Id))
//	return couchResp.StatusCode == http.StatusOK, couchResp
//}
