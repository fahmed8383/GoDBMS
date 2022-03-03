package Database

import (
	"errors"
)

// Create a global variable heap for this package. This variable stores the
// list of all tuples in memory.
var heap []*Tuple

// InitializeHeap is a function to initialize an empty heap list
func InitializeHeap() {
	heap = []*Tuple{}
}

// LoadHeap is a function to load a heap list from a heap list pointer
func LoadHeap(tuplesHeap *[]*Tuple) {
	heap = *tuplesHeap
}

// InsertTuple is a function to add a new tuple to the end of the heap
func InsertTuple(newTuple *Tuple) {
	heap = append(heap, newTuple)
}

// TupleExists is a function to check if a tuple exists in the heap given the key
// and the index of the key. If it exists return true, else return false.
func TupleExists(tupleKey interface{}, keyIndex int) bool {
	for _, currTuple := range heap {
		if tupleKey == interface{}(currTuple.Values[keyIndex]) {
			return true
		}
	}
	return false
}

// GetTuple is a function to get a tuple from the heap given the key and the
// index of the key. If the tuple exists return a pointer to it, otherwise
// return nil.
func GetTuple(tupleKey interface{}, keyIndex int) *Tuple {
	for _, currTuple := range heap {
		if tupleKey == currTuple.Values[keyIndex] {
			return currTuple
		}
	}

	return nil
}

// DeleteTuple is a function to remove a tuple from the heap given the key and
// the index of the key. If the tuple does not exist return an error stating so,
// otherwise return nil.
func DeleteTuple(tupleKey interface{}, keyIndex int) error {
	index := -1
	for i, currTuple := range heap {
		if tupleKey == currTuple.Values[keyIndex] {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("tuple does not exist")
	}

	heap = append(heap[:index], heap[index+1:]...)

	return nil
}

// GetHeap is a function to return a pointer to the heap.
func GetHeap() *[]*Tuple {
	return &heap
}
