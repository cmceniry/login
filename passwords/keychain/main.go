package main

import (
	"fmt"

	// tag::import[]
	keychain "github.com/keybase/go-keychain"
	// end::import[]
)

func main() {
	// tag::create[]
	err := keychain.AddItem(
		keychain.NewGenericPassword(
			";login example",
			"falken",
			";login example",
			[]byte("joshua"),
			"",
		),
	)
	// end::create[]
	// tag::duperr[]
	if err != nil && err != keychain.ErrorDuplicateItem {
		// end::duperr[]
		panic(err)
	}

	// tag::getgeneric[]
	item, err := keychain.GetGenericPassword(
		";login example",
		"falken",
		";login example",
		"",
	)
	// end::getgeneric[]
	if err != nil {
		panic(err)
	}
	// tag::printgeneric[]
	fmt.Println(string(item))
	// end::printgeneric[]

	// tag::query[]
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetLabel(";login example")
	query.SetAccount("falken")
	query.SetService(";login example")
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	// end::query[]
	if err != nil {
		panic(err)
	}
	// tag::printquery[]
	for _, i := range results {
		fmt.Println(string(i.Data))
	}
	// end::printquery[]

}
