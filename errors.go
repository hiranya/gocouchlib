package gocouchlib

import (
	"time"
	"fmt"
)

type CouchError struct {
	When	time.Time
	What	string
}

func (cerr *CouchError) Error() string {
	return fmt.Sprintf(">> [Exception] at %v, %s", cerr.When, cerr.What)
}


type GoCouchLibError struct {
	When	time.Time
	What	string
}

func (cerr *GoCouchLibError) Error() string {
	return fmt.Sprintf(">> [Exception] at %v, %s", cerr.When, cerr.What)
}