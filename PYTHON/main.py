import argparse
from fileParser.fileReader import fileToArray
from cards.blackjack import blackjack
import random

# Use of argparse for get the parameters of the command line.
# We define -i for intense mode, -np for nopick mode and 2 intengers for
# CardValue and count.
def getParameters():
    parser = argparse.ArgumentParser()
    parser.add_argument("-i","--intense", help="Use -i for intense mode: return probability from 1 to 8",
                        action="store_true")
    parser.add_argument("-np","--nopick", help="Use -np for nopick mode: probability without picking any card",
                        action="store_true")
    parser.add_argument('integers1', metavar='N', type=int,
                         help='an integer') 
    parser.add_argument('integers2', metavar='N', type=int,
                         help='an integer for the accu')                  
    args = parser.parse_args()
    if args.integers1 and args.integers2:
        global cardValue
        global count
        cardValue = args.integers1
        count = args.integers2
    if args.intense and args.nopick:
        print("Intense + nopick mode on")
        return 2
    if args.intense:
        print("Intense mode on")
        return 0
    elif args.nopick:
        print("Nopick mode on")
        return 1

if __name__ == '__main__':
    value = getParameters()
    nums = fileToArray("pruebita.txt")
    random.shuffle(nums)
    blackjack(nums,cardValue,count,value)

    
    