import time
from iterator.iterator import iterador

# blackjackProb - calculates the probability of not passing in blackjack, 
# we implement a Brute force algorithm that calculates the probabilities to have equal or 
# less than 21 
def blackjackProb(array,cardValue,count):
    start = time.time()
    wins=0
    plays=0
    a = iterador(len(array),count)
    while a.hasNext():
        hand = a.GetComb()
        result = cardValue
        for i in hand:
            result = result + array[i]
        if result  <= 21:
            wins = wins + 1
        plays = plays + 1

    
    porcentaje =  wins/plays * 100
    print("The probability of not passing in blackjack with a " + str(cardValue)  + 
    " choosing " + str(count) + " cards is of " + str(round(porcentaje, 6))+ 
    "%. TIME: "+ str(round((time.time() - start)*1000)) + "ms")

