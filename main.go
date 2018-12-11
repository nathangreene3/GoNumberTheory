package main

import (
	"fmt"
)

func main() {
	p := eratosthenes(1000)
	exportSequence(p, "test.csv")
	// p, err := importSequence("test.csv")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	fmt.Println(p)
}
