function iteratorClass () {
    let setSize;
    let combinationSize;
    let MAX_COMBINATIONS;
    let lastCombination;
    let completedCombinations;
    
    function init (m, n) {
        setSize = m;
        combinationSize = n;
        lastCombination = [...Array(n).keys()];
        lastCombination[n-1]--;
        completedCombinations = 0;
        MAX_COMBINATIONS = calulateCombinations(m, n);                
    }

    function hasNext () {
        return completedCombinations < MAX_COMBINATIONS;
    }

    function next () {
        let currentCombination = lastCombination;
        for (let i = combinationSize-1; i >= 0; i--) {
            if (currentCombination[i] < setSize - combinationSize + i) {
                currentCombination[i]++;
                break;
            } else {
                let j = i-1;
                while (currentCombination[j] >= setSize - combinationSize + j) {
                    j--;
                }
                currentCombination[j]++;                
                for (let k = j+1; k < combinationSize; k++) {
                    currentCombination[k] = currentCombination[k - 1] + 1;
                }
                break;
            }
        }
        completedCombinations++;
        lastCombination = currentCombination;
        return currentCombination;
    }

    

    return {
        init,
        hasNext,
        next
    }
}

function calulateCombinations (m, n) {
    let num = m;
    let denom = n;
    for(let i = m-1; i>m-n; i--) {
        num *= i;
    }
    for(let j = n-1; j>0; j--) {
        denom *= j 
    }
    return num/denom
}

module.exports = { iteratorClass, calulateCombinations };