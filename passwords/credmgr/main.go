package main

import (
	"fmt"

	// tag::import[]
	"github.com/danieljoos/wincred"
	// end::import[]
)

func main() {
	// tag::create[]
	cred := wincred.NewGenericCredential("loginExample")
	cred.UserName = "falken"
	cred.CredentialBlob = []byte("joshua")
	err := cred.Write()
	// end::create[]
	if err != nil {
		panic(err)
	}

	// tag::get[]
	cred, err := wincred.GetGenericCredential("loginExample")
	// end::get[]
	if err != nil {
		panic(err)
	}
	// tag::printget[]
	fmt.Println(cred.UserName)
	fmt.Println(strings(cred.CredentialBlob))
	// end::printget[]
}
