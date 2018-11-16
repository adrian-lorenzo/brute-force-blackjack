class iterador:

    # __init__ - Initializates the iterator
    # numberElements int - Number of elements to combinate with
    # combinationSize int - Size of each combination
    def __init__(self, numberElements, size):
        self.numberElements = numberElements
        self.size = size
        self.reset()

    # GetComb - Return an Array of ints representing one combination
    def GetComb(self):
        result = self.comb.copy()
        return result

    # NewComb - Creates a new combination
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

    # HasNext - returns true if there is another combination to get        
    def hasNext(self):
        return self.newComb()

    # Reset - Resets the iterator
    def reset(self):
        self.comb = []
        for i in range(0,self.size):
            self.comb.append(i)
        
        self.comb[self.size - 1] = self.comb[self.size - 1] - 1 



