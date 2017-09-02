// Echo3 prints the index and value of each command-line argument one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%v : %s\n", i, arg)
	}
}
