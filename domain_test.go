package main

import (
	"regexp"
	"testing"
)

func TestDomain(t *testing.T) {

	d := domain()
	// regex is "close enough"
	match, _ := regexp.MatchString("^[a-zA-Z0-9-]+\\.[a-zA-Z0-9-]+", d)

	if match == false {
		t.Errorf("Domain \"%q\" is invalid.", d)
	}
}
