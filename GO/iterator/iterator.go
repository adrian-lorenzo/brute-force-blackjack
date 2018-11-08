package iterator

// Iterator - for iterating through combinations
// comb: Last combination array
// numberElements: Number of elements to combinate with
// chosen: Size of each combination
type Iterator struct {
	comb                   []int
	numberElements, chosen int
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
	result := make([]int, iter.chosen)
	copy(result, iter.comb)
	return result
}

// NewComb - Creates a new combination
func (iter *Iterator) newComb() bool {
	last := iter.chosen - 1
	if iter.comb[last] != (iter.numberElements - iter.chosen + last) {
		iter.comb[last]++
	} else {
		i := last
		for iter.comb[i] == (iter.numberElements - iter.chosen + i) {
			if i <= 0 {
				return false
			}
			i--
		}
		iter.comb[i]++
		for j := i + 1; j < iter.chosen; j++ {
			iter.comb[j] = iter.comb[j-1] + 1
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
	for i := 0; i < iter.chosen; i++ {
		iter.comb = append(iter.comb, i)
	}
	iter.comb[iter.chosen-1]--
}
