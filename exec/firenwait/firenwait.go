// tag::setup[]
package main

import (
	"os"
	"os/exec"
)

func main() {
	// end::setup[]
	// tag::command[]
	c := exec.Command("/bin/ls", ".")
	// end::command[]
	// tag::connectoutput[]
	c.Stdout = os.Stdout
	// end::connectoutput[]
	// tag::run[]
	err := c.Run()
	if err != nil {
		panic(err)
	}
}

// end::run[]
