#!/usr/local/bin/python3

import argparse
import logging
import sys
from fileParser.fileReader import fileToArray
from bj.blackjackChecker import blackjackProb

cardValue = 0 # Value of the card selected
count = 0 # Number of cards to take

# Use of argparse for get the parameters of the command line.
# We define for get the int values , -intense for intense mode, -nopick for nopick mode 
# and 2 intengers for cardValue and count.
def getParameters():
    parser = argparse.ArgumentParser()
    parser.add_argument("-intense","--intense", help="Use -intense for intense mode: return probability from 1 to 8",
                        action="store_true")
    parser.add_argument("-nopick","--nopick", help="Use -nopick for nopick mode: probability without picking any card",
                        action="store_true")
    parser.add_argument('integers', metavar='Cards', type=int, nargs='*',
                         help='Number of card selected and/or number of cards to pick')                   
    args = parser.parse_args()

    global cardValue
    global count

    if args.intense and args.nopick:
        return 2
    if args.intense:
        if not args.integers:
            logging.exception("Not enough arguments")
            sys.exit(2)
        cardValue = args.integers[0]
        return 0
    elif args.nopick:
        if not args.integers:
            logging.exception("Not enough arguments")
            sys.exit(2)
        count = args.integers[0]
        return 1
    elif len(args.integers) < 2:
        logging.exception("Not enough arguments")
        sys.exit(2)
    cardValue = args.integers[0]
    count = args.integers[1]

# blackjack - checks which mode has to be run
# value is use as flag 
def execBlackjack(array,value):
    if value == 0:
        array.remove(cardValue)
        for i in range(1,9):
            blackjackProb(array,cardValue,i)
    elif value == 1:
        blackjackProb(array,0,count)
    elif value == 2:
        for i in range(1,9):
            blackjackProb(array,0,i)
    else:
        array.remove(cardValue)
        blackjackProb(array,cardValue,count)
    sys.exit(0)

if __name__ == '__main__':
    value = getParameters()
    nums = fileToArray("../deck.dat")
    execBlackjack(nums,value)
