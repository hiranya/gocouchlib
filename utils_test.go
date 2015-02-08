package gocouchlib

import (
	"fmt"
	"net/url"
	"testing"
)

func TestPut(t *testing.T) {
	var server = &Server{"http://couchdb1:5984",
		url.UserPassword("testuser", "password"),
	}
	db := Database{"gocouch", server}
	fmt.Println(db.Exists())

	if exists, _ := db.Exists(); !exists {
		t.Error("DB", db.DbName, "does not exists.")
	}
}
