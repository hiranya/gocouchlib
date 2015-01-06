package gocouchlib

import (
	"testing"
	"fmt"
	"net/url"
)

func TestDocumentExists(t *testing.T) {
	s := &Server{"http://localhost:5984",
		url.UserPassword("testuser", "password"),
	}

	db := &Database{"gocouch", s}
	
	doc1 := &Document{"doc1", "", db}
	
	exists, _ := doc1.Exists()
	fmt.Println("Doc1 exists:", exists)
}
