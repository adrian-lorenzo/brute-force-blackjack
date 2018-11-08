import random
from iterator.iterator import iterador

def blackjackProb(array,CardValue,count):
    wins=0
    plays=0
    
    a = iterador(len(array) ,count)

    while a.hasNext():
        hand = a.GetComb()
        result = CardValue
        for i in hand:
            result = result + array[i]
        if result  == 21:
            wins = wins + 1
        plays = plays + 1

    porcentaje =  wins/plays * 100

    print("La probabilidad de obtener 21 es: " + str(round(porcentaje, 2)) )