package main

import (
	"os"
	"os/exec"
	"time"
)

func main() {
	c := exec.Command("/bin/ls", ".")
	c.Stdout = os.Stdout
	// tag::start[]
	err := c.Start()
	// end::start[]
	if err != nil {
		panic(err)
	}
	// tag::wait[]
	go func() {
		err := c.Wait()
		if err != nil {
			panic(err)
		}
	}()
	// end::wait[]
	// tag::work[]
	// Do some other work...
	time.Sleep(2 * time.Second)
	// end::work[]
}
