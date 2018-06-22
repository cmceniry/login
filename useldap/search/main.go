package main

import (
	"fmt"
	"os"

	"github.com/cmceniry/login/useldap"
	ldap "gopkg.in/ldap.v2"
)

func main() {
	// tag::getpw[]
	pw, err := useldap.GetPassword()
	if err != nil {
		panic(err)
	}
	// end::getpw[]

	// tag::connect[]
	l, err := ldap.Dial("tcp", "localhost:389")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	// end::connect[]

	// tag::bind[]
	err = l.Bind("cn=admin,dc=example,dc=org", pw)
	if err != nil {
		panic(err)
	}
	// end::bind[]

	// tag::search[]
	results, err := l.Search(ldap.NewSearchRequest(
		"ou=people,dc=example,dc=org",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false,
		fmt.Sprintf("(cn=%s)", os.Args[1]),
		nil, nil,
	))
	// end::search[]
	if err != nil {
		panic(err)
	}
	// tag::show[]
	for _, r := range results.Entries {
		fmt.Printf("------- %s -------\n", r.DN)
		for _, attr := range r.Attributes {
			for _, v := range attr.Values {
				fmt.Printf("%s: %s\n", attr.Name, v)
			}
		}
	}
	// end::show[]
}
