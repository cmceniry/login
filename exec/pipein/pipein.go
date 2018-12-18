package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	// tag::command[]
	c := exec.Command("/usr/bin/dc")
	// end::command[]
	// tag::io[]
	c.Stdin = strings.NewReader("1\n2\n+\np\nq\n")
	c.Stdout = os.Stdout
	// end::io[]
	// tag::run[]
	err := c.Run()
	// end::run[]
	if err != nil {
		panic(err)
	}
}
