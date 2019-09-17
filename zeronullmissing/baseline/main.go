package main

import (
	"encoding/json"
	"fmt"
)

// tag::struct[]
type Items struct {
	FromZero    string
	FromNull    string
	FromPtrZero *string
	FromPtrNull *string
	FromMissing string
	FromPtrMiss *string
}

// end::struct[]

// tag::input[]
var input = `{
    "fromzero": "",
	"fromnull": null,
	"fromptrzero": "",
	"fromptrnull": null
}`

// end::input[]

func main() {
	// tag::output[]
	output := Items{}
	// end::output[]
	// tag::unmarshal[]
	err := json.Unmarshal([]byte(input), &output)
	// end::unmarshal[]
	if err != nil {
		panic(err)
	}
	// tag::print[]
	fmt.Printf("%#v\n", output)
	// end::print[]
}
