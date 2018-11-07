class iterador:

    def __init__(self, numberElements, size):
        self.numberElements = numberElements
        self.size = size
        self.reset()

    def GetComb(self):
        result = self.comb.copy()
        self.newComb()
        return result

    def newComb(self):
        for i in range(0,self.size):
            if self.comb[i] == (self.numberElements - self.size + i):
                if i == 0:
                    self.finished = True
                    break
                else:
                    self.comb[i-1] = self.comb[i-1] + 1
                    for j in range (i,self.size):
                        self.comb[j] = self.comb[j-1] + 1
                break
            elif i == (self.size -1 ):
                self.comb[i] = self.comb[i] + 1
                

    def HasNext(self):
        return not self.finished

    def reset(self):
        self.comb = []
        for i in range(0,self.size):
            self.comb.append(i)
        
        self.finished = False