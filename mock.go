package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	for i, a := range os.Args {
		fmt.Printf("%v) %v\n", i, a)
	}
}
