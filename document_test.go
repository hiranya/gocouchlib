package gocouchlib

import (
	"fmt"
	"net/url"
	//	"reflect"
	"testing"
)

func TestDocumentExists(t *testing.T) {
	s := &Server{"http://localhost:5984",
		url.UserPassword("testuser", "password"),
	}

	db := &Database{"gocouch", s}

	doc1 := &Document{Id: "doc1", Rev: "3-63cad646d83b402c86639c25d9dabd8a", Db: db}

	exists, couchResp := doc1.Exists()
	fmt.Println("Doc1 exists:", exists)
	fmt.Println("Doc1 content:", doc1, couchResp.StatusCode, couchResp.Headers)
}

func TestDocumentGet(t *testing.T) {
	s := &Server{"http://localhost:5984",
		url.UserPassword("testuser", "password"),
	}

	db := &Database{"gocouch", s}

	doc1 := &Document{Id: "doc1", Db: db}

	json, _ := doc1.Get()

	fmt.Println("=> TestDocumentGet():", json)

	switch json.(type) {
	case JsonObj:
	default:
		t.Error("Document.Get() did not return a Json document")
	}
}
