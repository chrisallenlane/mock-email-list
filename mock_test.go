package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

func TestMockPlaintext(t *testing.T) {

	// create a write buffer whose contents we can examine
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	// mock 5 emails
	mock(writer, 5, false)
	writer.Flush()

	// split the blob of text into an array of addresses
	got := strings.Split(strings.Trim(b.String(), "\n"), "\n") // is this right?

	// count the addresses returned
	i := 0

	// assert that each address is valid
	for _, email := range got {
		match, _ := regexp.MatchString("[a-zA-Z]{5}@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-]+", email)
		if match == false {
			t.Errorf("Email %q is invalid.", email)
		}
		i++
	}

	if i != 5 {
		t.Errorf("Wrong number of addresses returned: %q", i)
	}
}

func TestMockCSV(t *testing.T) {

	// create a write buffer whose contents we can examine
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	// mock 5 emails
	mock(writer, 5, true)
	writer.Flush()

	// split the blob of text into an array of addresses
	got := strings.Split(strings.Trim(b.String(), "\n"), "\n") // is this right?

	// count the addresses returned
	i := 0

	// assert that each address is valid
	for _, email := range got {
		match, _ := regexp.MatchString("[a-zA-Z]{5}@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-]+", email)
		if match == false {
			t.Errorf("Email %q is invalid.", email)
		}
		i++
	}

	if i != 5 {
		t.Errorf("Wrong number of addresses returned: %q", i)
	}
}

func BenchmarkMockPlaintext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mock(ioutil.Discard, 1000, false)
	}
}

func BenchmarkMockCSV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mock(ioutil.Discard, 1000, true)
	}
}
