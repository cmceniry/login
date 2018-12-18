package main

import (
	"os"
	"syscall"
)

func main() {
	// tag::env[]
	env := append(
		os.Environ(),
		"USENIXLOGIN=true",
	)
	// end::env[]
	// tag::handoff[]
	syscall.Exec("/bin/bash", []string{}, env)
	// end::handoff[]
}
