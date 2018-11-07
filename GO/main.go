package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/adrianlorenzodev/brute-force/go/iterator"
	"github.com/adrianlorenzodev/brute-force/go/parser"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("It's necessary more arguments")
	}

	data := parser.GetData(os.Args[1])
	num, errorAtoi := strconv.Atoi(os.Args[2])
	if errorAtoi != nil {
		log.Fatal(errorAtoi)
	}

	defer fmt.Printf("Done!\n")
	fmt.Printf("Start:\n")

	iter := iterator.InitIterator(len(data), num)

	start := time.Now().UnixNano() / int64(time.Millisecond)
	for iter.HasNext() {
		fmt.Printf("%v\n", iter.GetComb())
	}
	fmt.Printf("Time:  %d ms\n", (time.Now().UnixNano()/int64(time.Millisecond))-start)
}
