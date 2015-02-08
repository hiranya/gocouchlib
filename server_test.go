package gocouchlib

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {

	// Test case: CouchDB URL without authentication credentials
	s := &Server{"http://couchdb1:5984", nil}

	fmt.Println(s.FullUrl())

	if s.FullUrl() == "" {
		t.Error()
	}

	// Test case: CouchDB URL with authentication credentials
	s = &Server{"http://couchdb1:5984",
		url.UserPassword("testuser", "password"),
	}
	fmt.Println(s.FullUrl())

	if s.FullUrl() == "" || !strings.Contains(s.FullUrl(), "@") {
		t.Error()
	}

	// Test case: Server.Info()
	s = &Server{"http://couchdb1:5984",
		url.UserPassword("testuser", "password"),
	}

	fmt.Println(s.Info())

	// Test case: Server.AllDbs()
	allDbs := s.AllDbs()
	if len(allDbs.([]interface{})) == 0 {
		t.Error()
	}
	fmt.Println(allDbs)

}
