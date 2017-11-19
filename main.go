package main

import (
	"github.com/docopt/docopt-go"
	"os"
	"strconv"
)

func main() {

	// parse the command-line options
	opts, _ := docopt.Parse(usage(), nil, true, "1.0.0", false)

	// input conversions
	input := opts["<number-of-addresses>"].(string)
	number, err := strconv.Atoi(input)
	if err != nil {
		os.Stderr.WriteString("\"" + input + "\" must be a number.\n")
		os.Exit(1)
	}

	// generate email addresses
	mock(os.Stdout, number, opts["--csv"].(bool))
}
