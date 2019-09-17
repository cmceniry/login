package main

import (
	"encoding/json"
	"fmt"
)

var input = `{
    "fromzero": "",
	"fromnull": null,
	"fromptrzero": "",
	"fromptrnull": null
}`

func main() {
	// tag::decode[]
	output := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(input), &output)
	// end::decode[]
	if err != nil {
		panic(err)
	}
	// tag::loop[]
	keys := []string{"fromzero", "fromnull", "fromptrzero",
		"fromptrnull", "frommissing"}
	for _, k := range keys {
		// end::loop[]
		// tag::check[]
		v, ok := output[k]
		if !ok {
			fmt.Printf(`"%s" is missing`+"\n", k)
			continue
		}
		// end::check[]
		// tag::type[]
		switch v.(type) {
		case nil:
			fmt.Printf(`"%s" is null`+"\n", k)
		case string:
			fmt.Printf(`"%s" is present and equal to "%s"`+"\n", k, v.(string))
		default:
			fmt.Printf(`"%s" unhandled type %T`+"\n", k, v)
		}
		// end::type[]
	}
}
