package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	// tag::vars[]
	feed := []string{"a", "b", "c", "d", "e", ""}
	c := exec.Command("/bin/cat")
	// end::vars[]

	// tag::stdin[]
	cin, err := c.StdinPipe()
	// end::stdin[]
	if err != nil {
		panic(err)
	}
	// tag::stdout[]
	cout, err := c.StdoutPipe()
	// end::stdout[]
	if err != nil {
		panic(err)
	}

	// tag::buffer[]
	bin := bufio.NewWriter(cin)
	bout := bufio.NewScanner(cout)
	// end::buffer[]

	// tag::prime[]
	bin.WriteString("\n")
	bin.Flush()
	c.Start()
	// end::prime[]

	// tag::addnprint[]
	for _, addition := range feed {
		if !bout.Scan() {
			panic("ended early")
		}
		if bout.Text() != "" {
			fmt.Printf("%s\n", bout.Text())
		}
		bin.WriteString(bout.Text() + addition + "\n")
		bin.Flush()
	}
	// end::addnprint[]

	// tag::cleanup[]
	cin.Close()
	c.Wait()
	// end::cleanup[]
}
