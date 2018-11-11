import random
from iterator.iterator import iterador
from timeit import default_timer as timer

def blackjack(array,CardValue,count,value):
    if value == 0:
        for i in range(2,8):
            blackjackProb(array,CardValue,i)
    elif value == 1:
        blackjackProb(array,0,count)
    elif value == 2:
        for i in range(2,8):
            blackjackProb(array,0,i)
    else:
        blackjackProb(array,CardValue,count)

def blackjackProb(array,CardValue,count):
    start = timer()
    wins=0
    plays=0
    
    a = iterador(len(array) ,count)

    while a.hasNext():
        hand = a.GetComb()
        result = CardValue
        for i in hand:
            result = result + array[i]
        if result  <= 21:
            wins = wins + 1
        plays = plays + 1

    porcentaje =  wins/plays * 100
    end = timer()
    print("La probabilidad de obtener 21 es: " + str(round(porcentaje, 2)) + "%"
    "\nEl tiempo de ejecucion del procedimiento es:" + str(round(end - start, 4)) + " segundos")
    