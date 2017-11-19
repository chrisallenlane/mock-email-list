package main

import (
	"fmt"
	"io"
)

func mock(writer io.Writer, number int, csv bool) {

	// keep the conditional test outside of the loop for performance
	if csv == true {
		for i := 1; i <= number; i++ {
			fmt.Fprintf(writer, "\""+username(10)+"@"+domain()+"\"\n")
		}
	} else {
		for i := 1; i <= number; i++ {
			fmt.Fprintf(writer, username(10)+"@"+domain()+"\n")
		}
	}
}
