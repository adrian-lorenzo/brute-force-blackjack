from timeit import default_timer as timer
from fileParser.fileReader import fileToArray
from cards.blackjack import blackjackProb
import random

if __name__ == '__main__':
    nums = fileToArray("pruebita.txt")
    random.shuffle(nums)
    blackjackProb(nums,10,2)