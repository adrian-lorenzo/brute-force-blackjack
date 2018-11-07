from iterator import iterador
from timeit import default_timer as timer

if __name__ == '__main__':
    a = iterador(50,10)
    start = timer()
    while a.HasNext():
        print(a.GetComb())
    end = timer()
    print(end - start)