package main

import (
	"fmt"
)

func main() {
	n := 100
	list := numberDivisorList(n)
	for i := range list {
		fmt.Printf("list[%d] = %d\n", i, *list[i])
	}
}
