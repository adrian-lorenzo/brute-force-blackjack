package iterator

// Iterator - for iterating through combinations
// comb: Last combination array
// numberElements: Number of elements to combinate with
// chosen: Size of each combination
type Iterator struct {
	comb                   []int
	numberElements, chosen int
	finished               bool
}

// InitIterator - Initilization of the iterator
func InitIterator(numberElements, chosen int) *Iterator {
	iter := Iterator{}
	iter.numberElements = numberElements
	iter.chosen = chosen
	iter.Reset()
	return &iter
}

// GetComb - Return a []int representing one combination
func (iter *Iterator) GetComb() []int {
	defer iter.newComb()
	result := make([]int, iter.chosen)
	copy(result, iter.comb)
	return result
}

// NewComb - Creates a new combination
func (iter *Iterator) newComb() {

	for i := 0; i < iter.chosen; i++ {
		if iter.comb[i] == (iter.numberElements - iter.chosen + i) {
			if i == 0 {
				iter.finished = true
				break
			} else {
				iter.comb[i-1]++
				for j := i; j < iter.chosen; j++ {
					iter.comb[j] = iter.comb[j-1] + 1
				}
			}
			break
		} else if i == (iter.chosen - 1) {
			iter.comb[i]++
		}
	}

}

// HasNext - returns true if there is another combination to get
func (iter *Iterator) HasNext() bool {
	return !iter.finished
}

// Reset - Resets the iterator
func (iter *Iterator) Reset() {
	iter.comb = make([]int, iter.chosen)
	for i := 0; i < iter.chosen; i++ {
		iter.comb[i] = i
	}

	iter.finished = false
}
