package gocouchlib

import (
    "testing"
    "time"
    "fmt"
)

func TestErrors(t *testing.T) {
	err := run()
	
	fmt.Println(err)
	
	if err == nil {
		t.Error()
	}
}

func run() error {
	return &CouchError{
		time.Now(), "Test error thrown from CouchDB",
	}
}

