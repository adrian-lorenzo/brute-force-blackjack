module.exports = class {
    
    constructor(m,n) {
        this.init(m,n)
    }
    
    init (m, n) {
        this.setSize = m;
        this.isNext = true;
        this.combinationSize = n;
        this.currentCombination = [...Array(n).keys()];
        this.currentCombination[n-1]--;
    }

    hasNext () {
        return this.isNext;
    }

    next () {
        let last = this.combinationSize-1
        if (this.currentCombination[last] < this.setSize - this.combinationSize + last) {
            this.currentCombination[last]++;
        } else {
            let j = last;
            while (this.currentCombination[j] === this.setSize - this.combinationSize + j) {
                if (j <= 0) {
                    this.isNext = false;
                    return this.currentCombination;
                }
                j--;
            }
            this.currentCombination[j]++;                
            for (let k = j+1; k < this.combinationSize; k++) {
                this.currentCombination[k] = this.currentCombination[k - 1] + 1;
            }
        }
        return this.currentCombination;
    }
}

