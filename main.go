package main

import (
	"fmt"
	api "xeno_go/api.go/xeno_go"
)


func main() {
	results := api.Query{
		Search: "",
		Parameters: map[string]string{
			"area": "america",
			"since": "50",
		},
	}.GetRecordings()
	fmt.Printf("\nResults found: %v", len(results))
}
