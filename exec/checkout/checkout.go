package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

func main() {
	// tag::command[]
	c := exec.Command("/usr/bin/false")
	err := c.Run()
	// end::command[]
	// tag::result[]
	switch err.(type) {
	case *exec.ExitError:
		ws := c.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Printf("Exited %d\n", ws.ExitStatus())
	case nil:
		fmt.Printf("Exited normally\n")
	default:
		panic(err)
	}
	// end::result[]
}
