function iteratorClass () {
    let setSize;
    let combinationSize;
    let MAX_COMBINATIONS;
    let currentCombination;
    let completedCombinations;
    
    function init (m, n) {
        setSize = m;
        combinationSize = n;
        currentCombination = [...Array(n).keys()];
        currentCombination[n-1]--;
        completedCombinations = 0;
        MAX_COMBINATIONS = calulateCombinations(m, n);                
    }

    function hasNext () {
        return completedCombinations < MAX_COMBINATIONS;
    }

    function next () {
        let i = combinationSize-1
        if (currentCombination[i] < setSize - combinationSize + i) {
            currentCombination[i]++;
        } else {
            let j = i-1;
            while (currentCombination[j] >= setSize - combinationSize + j) {
                j--;
            }
            currentCombination[j]++;                
            for (let k = j+1; k < combinationSize; k++) {
                currentCombination[k] = currentCombination[k - 1] + 1;
            }
        }
        completedCombinations++;
        return currentCombination;
    }

    function aNext() {
        for (let i = 0; i < combinationSize; i++) {
            if (currentCombination[i] === (setSize - combinationSize + i)) {
                currentCombination[i-1]++
                for (let j = i; j < combinationSize; j++) {
                    currentCombination[j] = currentCombination[j-1] + 1;
                }
                break;
            } else if (i == (combinationSize - 1)) {
                currentCombination[i]++
            }
        }
        completedCombinations++;
        return currentCombination
    }



    return {
        init,
        hasNext,
        aNext,
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