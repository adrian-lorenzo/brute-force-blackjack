package main

import (
	"fmt"
	"time"

	"github.com/adrianlorenzodev/brute-force/GO/iterator"
	_ "github.com/adrianlorenzodev/brute-force/GO/parser"
)

func main() {
	iter := iterator.InitIterator(100, 95)
	start := time.Now().UnixNano() / int64(time.Millisecond)

	defer fmt.Printf("Done!\n")
	fmt.Printf("Start:\n")
	for iter.HasNext() {
		iter.GetComb()
	}
	fmt.Printf("Time:  %d ms\n", (time.Now().UnixNano()/int64(time.Millisecond))-start)
}
