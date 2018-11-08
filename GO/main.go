package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/adrianlorenzodev/brute-force/go/iterator"
	"github.com/adrianlorenzodev/brute-force/go/parser"
)

func main() {
	help := flag.Bool("help", false, "Print help about the command")
	intense := flag.Bool("intense", false, "Print all probabilties from 1 to 10 cards")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		if *help {
			fmt.Printf("brute-force <Card> <Number of cards>\n\n\n")
			os.Exit(2)
		} else if len(args) < 1 || (len(args) == 1 && !*intense) {
			log.Fatal("It's necessary more arguments")
		}
	}

	cardSelected, errorAtoi := strconv.Atoi(args[0])
	if errorAtoi != nil {
		log.Fatal(errorAtoi)
	}

	cards := parser.Shuffle(parser.GetData("deck.dat"))
	for i, card := range cards {
		if card == cardSelected {
			cards = append(cards[:i], cards[i+1:]...)
			break
		}
	}

	if *intense {
		for i := 1; i < 11; i++ {
			probabilityCheck(cards, cardSelected, i)
		}
	} else {
		num, errorAtoi := strconv.Atoi(args[1])
		if errorAtoi != nil {
			log.Fatal(errorAtoi)
		}
		probabilityCheck(cards, cardSelected, num)
	}

}

func probabilityCheck(cards []int, cardSelected, num int) {
	iter := iterator.InitIterator(len(cards), num)
	start := time.Now().UnixNano() / int64(time.Millisecond)

	tries := 0.00
	wins := 0.00

	for iter.HasNext() {
		comb := iter.GetComb()
		result := 0
		for _, number := range comb {
			result += cards[number]
		}

		if (result + cardSelected) <= 21 {
			wins++
		}

		tries++
	}

	fmt.Printf("The probability of not passing in blackjack with a %d choosing %d cards is of %f%%. TIME: %dms\n", cardSelected, num,
		((wins / tries) * 100), (time.Now().UnixNano()/int64(time.Millisecond))-start)
}
