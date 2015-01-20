package gocouchlib

import (
	"fmt"
	"net/url"
	"testing"
)

type EmployeeDoc struct {
	EmployeeId   int
	EmployeeName string
	EmployeeAge  int
}

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

func TestDocumentSave(t *testing.T) {
	s := &Server{"http://localhost:5984",
		url.UserPassword("testuser", "password"),
	}

	db := &Database{"gocouch", s}

	// case: both _id and _rev are not specified
	new_doc1 := &Document{Db: db, Json: EmployeeDoc{10, "Hiranya", 32}}
	fmt.Println("=> new_doc1: ", new_doc1)
	success, couchResp := new_doc1.Save()
	if !success {
		t.Error("New document without _id and _rev did not save successfully. CouchResponse:", couchResp)
	}

	// case: only _id is specified, no _rev
	new_doc2 := &Document{Db: db, Id: "new_doc2", Json: EmployeeDoc{20, "John", 12}}
	success, couchResp = new_doc2.Save()
	if !success {
		t.Error("New document with _id specified but without _rev, did not save successfully. CouchResponse:", couchResp)
	}

	success, couchResp = new_doc2.Delete()
	if !success {
		t.Error("new_doc2 deletion not successful. CouchResponse:", couchResp)
	}

	// case: both _id and _rev is specified. Therefore an update to doc
	new_doc3 := &Document{Db: db, Id: "new_doc3", Json: EmployeeDoc{30, "Deshani", 22}}
	success, couchResp = new_doc3.Save()
	if !success {
		t.Error("new_doc3 did not save successfully. CouchResponse:", couchResp)
	}

	emp2 := new_doc3.Json.(EmployeeDoc)
	emp2.EmployeeName = "Chirag"
	new_doc3.Json = emp2
	success, couchResp = new_doc3.Save()
	if !success {
		t.Error("new_doc3 second save was not successful. CouchResponse:", couchResp)
	}

	success, couchResp = new_doc3.Delete()
	if !success {
		t.Error("new_doc3 deletion was not successful. CouchResponse:", couchResp)
	}
}

func TestDocumentDelete(t *testing.T) {
	s := &Server{"http://localhost:5984",
		url.UserPassword("testuser", "password"),
	}

	db := &Database{"gocouch", s}

	new_delete_doc1 := &Document{Id: "new_delete_doc1", Db: db, Json: EmployeeDoc{100, "Hiranya", 33}}
	fmt.Println("=> new_delete_doc1: ", new_delete_doc1)
	success, couchResp := new_delete_doc1.Save()
	if !success {
		t.Error("new_delete_doc1 did not save successfully. CouchResponse:", couchResp)
	}

	success, couchResp = new_delete_doc1.Delete()
		fmt.Println("=> new_delete_doc1: ", new_delete_doc1)

	if !success {
		t.Error("new_delete_doc1 deletion was not successful. CouchResponse:", couchResp)
	}
	
	if !new_delete_doc1.Deleted {
		t.Error("new_delete_doc1's deleted attribute (corresponds to _deleted) has not been set to TRUE")
	}
}
