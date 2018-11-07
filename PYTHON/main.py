from timeit import default_timer as timer
<<<<<<< HEAD
from iterator.iterator import iterador
=======
from iterator import iterador
>>>>>>> 271398996e9fe3203e6a0f70799242f93770ef25

if __name__ == '__main__':
    a = iterador(20,10)
    start = timer()
    while a.HasNext():
        print(a.GetComb())
    end = timer()
    print(end - start)
