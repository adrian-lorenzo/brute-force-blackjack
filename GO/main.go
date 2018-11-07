package main

import (
	"fmt"
	"github.com/adrianlorenzodev/brute-force/GO/iterator"
)

func main() {
	iter := iterator.InitIterator(10, 3)
	i := 0
	for iter.HasNext() {
		fmt.Printf("%d: %v \n", i, iter.GetComb())
		i++
	}
}
