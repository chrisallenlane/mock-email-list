package main

func usage() string {

	return `mock-email-list

Usage:
  mock-email-list <number-of-addresses> [--csv]

Options:
  -h --help     Show this screen.
  -v --version  Show version.
  -c --csv      Output CSV.

Examples:

  To mock 1000 email addresses:
    mock-email-list 1000 > emails.txt`
}
