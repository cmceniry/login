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
	// end::getpw[]
	if err != nil {
		panic(err)
	}

	// tag::connect[]
	l, err := ldap.Dial("tcp", "localhost:389")
	// end::connect[]
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer l.Close()

	// tag::bind[]
	err = l.Bind("cn=admin,dc=example,dc=org", pw)
	// end::bind[]
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// tag::search[]
	results, err := l.Search(ldap.NewSearchRequest(
		"ou=groups,dc=example,dc=org",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false,
		"(cn=mygroup)",
		nil, nil,
	))
	// end::search[]
	if err != nil {
		panic(err)
	}
	if len(results.Entries) == 0 {
		panic("group not found")
	}
	// tag::exist[]
	members := results.Entries[0].GetAttributeValues("member")
	for _, v := range members {
		if v == os.Args[1] {
			os.Exit(0)
		}
	}
	// end::exist[]

	// tag::modify[]
	m := ldap.NewModifyRequest(
		"cn=mygroup,ou=groups,dc=example,dc=org",
	)
	m.Add("member", []string{os.Args[1]})
	err = l.Modify(m)
	if err != nil {
		panic(err)
	}
	// end::modify[]

}
