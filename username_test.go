package main

import (
	"regexp"
	"testing"
)

func TestUsername(t *testing.T) {

	user := username(10)
	match, _ := regexp.MatchString("^[a-zA-Z]{10}$", user)

	if match == false {
		t.Errorf("Username \"%q\" is invalid.", user)
	}
}
