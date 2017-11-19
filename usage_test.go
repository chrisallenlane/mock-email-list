package main

import (
	"regexp"
	"testing"
)

func TestUsage(t *testing.T) {

	text := usage()
	match, _ := regexp.MatchString("^mock-email-list", text)

	if match == false {
		t.Errorf("Text returned by usage() is invalid.")
	}
}
