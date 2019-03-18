package main

import (
	"fmt"

	// tag::import[]
	"github.com/99designs/keyring"
	// end::import[]
)

func main() {
	// tag::open[]
	kr, err := keyring.Open(keyring.Config{
		ServiceName: ";login example",
	})
	// end::open[]
	if err != nil {
		panic(err)
	}

	// tag::create[]
	_ = kr.Set(keyring.Item{
		Key:  "falken",
		Data: []byte("joshua"),
	})
	// end::create[]

	// tag::get[]
	p, err := kr.Get("falken")
	// end::get[]
	if err != nil {
		panic(err)
	}
	// tag::print[]
	fmt.Println(string(p.Data))
	// end::print[]
}
