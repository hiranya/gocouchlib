package gocouchlib

import (
	"net/http"
)

type Database struct {
	DbName string
	Server *Server
}

func (db *Database) endpoint(api string) string {
	return db.Server.FullUrl() + "/" + db.DbName + api
}

// Check whether the database exists
// HEAD /{db}
func (db *Database) Exists() (bool, *CouchResponse) {
	couchResp, _ := httpClient.Head(db.endpoint("/"), nil)
	return couchResp.StatusCode == http.StatusOK, couchResp
}

// Retrieves database information
// GET /{db}
func (db *Database) Info() (JsonObj, *CouchResponse) {
	couchResp, _ := httpClient.Get(db.endpoint("/"), nil)
	return couchResp.Json, couchResp
}

func (db *Database) Create() (bool, *CouchResponse) {
	couchResp, _ := httpClient.Put(db.endpoint("/"))
	return couchResp.StatusCode == http.StatusCreated, couchResp
}

// Deletes database
// DELETE /{db}
func (db *Database) Delete() (bool, *CouchResponse) {
	couchResp, _ := httpClient.Delete(db.endpoint("/"))
	return couchResp.StatusCode == http.StatusOK, couchResp
}
