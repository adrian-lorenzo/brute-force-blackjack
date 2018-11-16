import argparse
from fileParser.fileReader import fileToArray
from bj.blackjackChecker import blackjack
import random
import logging
import sys

# getParameters - Use of argparse for get the parameters of the command line.
# We define -i for intense mode, -np for nopick mode 
# and intengers for CardValue and count.
def getParameters():
    parser = argparse.ArgumentParser()
    parser.add_argument("-i","--intense", help="Use -i for intense mode: return probability from 1 to 8",
                        action="store_true")
    parser.add_argument("-np","--nopick", help="Use -np for nopick mode: probability without picking any card",
                        action="store_true")
    parser.add_argument('integers', metavar='N', type=int, nargs='*',
                         help='an integer')                   
    args = parser.parse_args()
    global cardValue
    global count
    if args.intense and args.nopick:
        cardValue = 0
        count = 0
        return 2
    if args.intense:
        if not args.integers:
            logging.exception("Not enough arguments")
            sys.exit(2)
        count = 0
        cardValue = args.integers[0]
        return 0
    elif args.nopick:
        if not args.integers:
            logging.exception("Not enough arguments")
            sys.exit(2)
        cardValue = 0
        count = args.integers[0]
        return 1
    elif len(args.integers) < 2:
        logging.exception("Not enough arguments")
        sys.exit(2)
    cardValue = args.integers[0]
    count = args.integers[1]

if __name__ == '__main__':
    value = getParameters()
    nums = fileToArray("../deck.dat")
    random.shuffle(nums)
    blackjack(nums,cardValue,count,value)
    sys.exit(0)

    
