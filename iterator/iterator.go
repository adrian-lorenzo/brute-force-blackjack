package iterators

// Iterator - for iterating through combinations
// comb: Last combination array
// numberElements: Number of elements to combinate with
// chosen: Size of each combination
// indexIterate: Index of the array where the data is being altered
// max: Max value that the value at indexInterate position will have
type Iterator struct {
	comb                                      []int
	indexIterate, max, numberElements, chosen int
}

// Initilization of the iterator
func (iter *Iterator) initIterator(numberElements, chosen int) {
	iter.numberElements = numberElements
	iter.chosen = chosen
	iter.reset()
}

// Return a []int representing one combination
func (iter *Iterator) getComb() []int {
	defer iter.newComb()
	return iter.comb
}

// Creates a new combination
func (iter *Iterator) newComb() bool {
	if iter.indexIterate == -1 {
		return false
	}

	iter.comb[iter.indexIterate]++
	if iter.max == iter.comb[iter.indexIterate] {
		iter.max--
		iter.indexIterate--
	}

	return true
}

func (iter *Iterator) hasNext() bool {
	if iter.indexIterate == -1 {
		return false
	}

	return true
}

// Resets the iterator
func (iter *Iterator) reset() {
	for i := 0; i < iter.chosen; i++ {
		iter.comb[i] = i
	}

	iter.indexIterate = iter.chosen
	iter.max = iter.numberElements
}
