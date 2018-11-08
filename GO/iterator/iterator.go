package iterator

// Iterator - for iterating through combinations
// combination []int - Last combination array
// numberElements int - Number of elements to combinate with
// combinationSize int - Size of each combination
type Iterator struct {
	combination                     []int
	numberElements, combinationSize int
}

// InitIterator - Initilization of the iterator
// numberElements int - numberElements to choose from when combinating
// combinationSize int - size of the combinations
func InitIterator(numberElements, combinationSize int) *Iterator {
	iter := Iterator{}
	iter.numberElements = numberElements
	iter.combinationSize = combinationSize
	iter.Reset()
	return &iter
}

// Next - Return a []int representing one combination
func (iter *Iterator) Next() []int {
	result := make([]int, iter.combinationSize)
	copy(result, iter.combination)
	return result
}

// NewComb - Creates a new combination
func (iter *Iterator) newComb() bool {
	last := iter.combinationSize - 1
	if iter.combination[last] != (iter.numberElements - iter.combinationSize + last) {
		iter.combination[last]++
	} else {
		i := last
		for iter.combination[i] == (iter.numberElements - iter.combinationSize + i) {
			if i <= 0 {
				return false
			}
			i--
		}
		iter.combination[i]++
		for j := i + 1; j < iter.combinationSize; j++ {
			iter.combination[j] = iter.combination[j-1] + 1
		}
	}
	return true
}

// HasNext - returns true if there is another combination to get
func (iter *Iterator) HasNext() bool {
	return iter.newComb()
}

// Reset - Resets the iterator
func (iter *Iterator) Reset() {
	for i := 0; i < iter.combinationSize; i++ {
		iter.combination = append(iter.combination, i)
	}
	iter.combination[iter.combinationSize-1]--
}
