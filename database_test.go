package gocouchlib

import (
	"fmt"
	"net/url"
	"testing"
)

func TestDatabaseCreateDelete(t *testing.T) {
	s := &Server{"http://localhost:5984",
		url.UserPassword("testuser", "password"),
	}

	db := &Database{"gocouch3", s}

	isCreated, _ := db.Create()
	fmt.Println("isCreated:", db.DbName, isCreated)

	if !isCreated {
		t.Error("DB creation failed")
	}

	// 	Test case: Database.Delete()
	isDeleted, _ := db.Delete()
	fmt.Println("isDeleted:", db.DbName, isDeleted)
}

func TestDatabaseExists(t *testing.T) {
	// Test case: Database.Exists()
	s := &Server{"http://localhost:5984",
		url.UserPassword("testuser", "password"),
	}

	db := &Database{"gocouch", s}

	if exists, _ := db.Exists(); !exists {
		t.Error()
	}

	//Test case: Database.DbInfo()
	dbInfo, _ := db.Info()
	dbName := dbInfo.(map[string]interface{})["db_name"]
	if dbName != "gocouch" {
		t.Error()
	}
	fmt.Println("[Log]: DbName:", dbName)

}
