class iterador:

    def __init__(self, numberElements, size):
        self.numberElements = numberElements
        self.size = size
        self.reset()

    def GetComb(self):
        result = self.comb.copy()
        return result

    def newComb(self):
        last = self.size - 1
        if self.comb [last] < (self.numberElements - self.size + last):
            self.comb[last] = self.comb[last] + 1
        else:

            j = last
            while self.comb[j] == (self.numberElements - self.size + j ):
                if j == 0:
                    return False
                j = j - 1
            self.comb[j] = self.comb[j] + 1   
            for k in range(j + 1,self.size):
                self.comb[k] = self.comb[k-1] + 1
        return True
            
    def hasNext(self):
        return self.newComb()

    def reset(self):
        self.comb = []
        for i in range(0,self.size):
            self.comb.append(i)
        
        self.comb[self.size - 1] = self.comb[self.size - 1] - 1 



