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

var (
	help, intense, nopick *bool //flags for options
	args                  []string
)

func init() {
	help = flag.Bool("help", false, "Print help about the command")
	intense = flag.Bool("intense", false, "Print all probabilties from 1 to 8 cards")
	nopick = flag.Bool("nopick", false, "Print probability without picking any card")

	flag.Parse()
	args = flag.Args()
}

func main() {
	//Command completion checker
	if len(args) < 2 {
		if *help {
			fmt.Printf("\n\n.\\blackjack <num-first-card> <num-to-take>/-intense" +
				"\n\nCalculates the probability of not passing if you have one card when taking more in blackjack." +
				"\n\nOptions:" +
				"\n\n-help           Prints information about the command." +
				"\n-intense        Generates the probability of not passing from 1 to 8 cards more taken." +
				"\n-nopick         The command ignores <num-first-card> argument and calculates the probability" +
				"\n                without taking any card. It works with -intense option.\n\n")
			os.Exit(2)
		} else if (len(args) == 1 && !*intense && !*nopick) || (len(args) == 0 && (!*intense || !*nopick)) {
			log.Fatal("It's necessary more arguments")
		}
	}

	cards := parser.GetNumbers("deck.dat")
	var cardSelected, numCards, numArg int
	var errorAtoi error

	if !*nopick {
		cardSelected, errorAtoi = strconv.Atoi(args[0])
		if errorAtoi != nil {
			log.Fatal(errorAtoi)
		}
		for i, card := range cards {
			if card == cardSelected {
				cards = append(cards[:i], cards[i+1:]...)
				break
			}
		}
		numArg = 1
	} else {
		cardSelected = 0
		numArg = 0
	}

	numCards, errorAtoi = strconv.Atoi(args[numArg])
	if errorAtoi != nil {
		log.Fatal(errorAtoi)
	}

	blackJack(cards, cardSelected, numCards, *intense)

}

func blackJack(cards []int, cardSelected, numCards int, intense bool) {
	if intense {
		for i := 1; i < 9; i++ {
			probabilityCheck(cards, cardSelected, i)
		}
	} else {
		probabilityCheck(cards, cardSelected, numCards)
	}
}

func probabilityCheck(cards []int, cardSelected, numCards int) {
	iter := iterator.InitIterator(len(cards), numCards)

	tries := 0.00
	wins := 0.00

	start := time.Now().UnixNano() / int64(time.Millisecond)

	for iter.HasNext() {
		comb := iter.Next()
		result := 0
		for _, number := range comb {
			result += cards[number]
		}

		if (result + cardSelected) <= 21 {
			wins++
		}

		tries++
	}

	fmt.Printf("The probability of not passing in blackjack with a %d choosing %d cards is of %f%%. TIME: %dms\n",
		cardSelected, numCards, ((wins / tries) * 100), (time.Now().UnixNano()/int64(time.Millisecond))-start)
}
