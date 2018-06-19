package useldap

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// GetPassword is a simple function to obtain a password from the
// command line without echoing it
//
// tag::getpassword[]
func GetPassword() (string, error) {
	fmt.Printf("Password: ")
	pw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(pw), nil
}

// end::getpassword[]
