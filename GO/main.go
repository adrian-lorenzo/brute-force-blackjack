package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/adrianlorenzodev/brute-force/go/bj"
	"github.com/adrianlorenzodev/brute-force/go/parser"
)

var (
	help, intense, nopick *bool //flags for options
	args                  []string
	numArg                int
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
			helpOutput()
		} else if (len(args) == 1 && !*intense && !*nopick) ||
			(len(args) == 0 && (!*intense || !*nopick)) {
			log.Fatal("It's necessary more arguments")
		}
	}

	cards := parser.GetNumbers("../deck.dat")
	var cardSelected, numCards int
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

	if *intense {
		for i := 1; i < 9; i++ {
			blackjack.ProbabilityCheck(cards, cardSelected, i)
		}
	} else {
		numCards, errorAtoi = strconv.Atoi(args[numArg])
		if errorAtoi != nil {
			log.Fatal(errorAtoi)
		}
		blackjack.ProbabilityCheck(cards, cardSelected, numCards)
	}
}

// helpOutput - Prints the output of the -help option
func helpOutput() {

	message := `
.\blackjack [-help] [-intense [FIRST_CARD]] [-nopick [NUMBER OF DRAWS]] [FIRST_CARD] [NUMBER OF DRAWS]

Calculates the probability of not passing if you have one card when taking more in blackjack.
				
Options:
				
-help           Prints information about the command.
-intense        Generates the probability of not passing from 1 to 8 cards more taken.
-nopick         The command ignores [FIRST_CARD] argument and calculates the probability
		without taking any card. It works with -intense option.
									
`

	fmt.Printf(message)
	os.Exit(2)
}
