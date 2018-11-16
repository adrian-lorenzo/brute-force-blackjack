import sys
from fileParser.fileReader import fileToArray
from bj.blackjack import blackjack
import random

def getParameters(parameters):
    if len(parameters) > 1:
        if parameters[1] == "--help":
            print("--Use -i for intense mode: return probability from 1 to 11\n" + 
            "--No parameters for normal use")
            return -1
        if parameters[1] == "-i":
            return 1
        else: 
            print("--Use --help to get information about the command")
            return -1

if __name__ == '__main__':
    value = getParameters(sys.argv)
    nums = fileToArray("../deck.dat")
    random.shuffle(nums)
    blackjack(nums,5,2,value)