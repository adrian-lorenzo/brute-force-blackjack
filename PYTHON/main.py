from timeit import default_timer as timer
from iterator2 import iterador

if __name__ == '__main__':
    a = iterador(20,10)
    start = timer()
    while a.HasNext():
        print(a.GetComb())
    end = timer()
    print(end - start)