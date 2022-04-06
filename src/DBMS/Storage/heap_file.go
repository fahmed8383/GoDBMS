package Storage

import (
	"errors"
)

// Create a global variable heap for this package. This variable stores the
// list of all tuples in memory.
type Heap struct {
	Tuples []*Tuple
}

// InitializeHeap is a function to initialize an empty heap list
func InitializeHeap() (*Heap) {
	return &Heap{[]*Tuple{}}
}

// InsertTuple is a function to add a new tuple to the end of the heap
func (heap *Heap) InsertTuple(newTuple *Tuple) {
	heap.Tuples = append(heap.Tuples, newTuple)
}

// TupleExists is a function to check if a tuple exists in the heap given the key
// and the index of the key. If it exists return true, else return false.
func (heap *Heap) TupleExists(tupleKey interface{}, keyIndex int) bool {
	for _, currTuple := range heap.Tuples {
		if tupleKey == interface{}(currTuple.Values[keyIndex]) {
			return true
		}
	}
	return false
}

// GetTuple is a function to get a tuple from the heap given the key and the
// index of the key. If the tuple exists return a pointer to it, otherwise
// return nil.
func (heap *Heap) GetTuple(tupleKey interface{}, keyIndex int) *Tuple {
	for _, currTuple := range heap.Tuples {
		if tupleKey == currTuple.Values[keyIndex] {
			return currTuple
		}
	}

	return nil
}

// DeleteTuple is a function to remove a tuple from the heap given the key and
// the index of the key. If the tuple does not exist return an error stating so,
// otherwise return nil.
func (heap *Heap) DeleteTuple(tupleKey interface{}, keyIndex int) error {
	index := -1
	for i, currTuple := range heap.Tuples {
		if tupleKey == currTuple.Values[keyIndex] {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("tuple does not exist")
	}

	heap.Tuples = append(heap.Tuples[:index], heap.Tuples[index+1:]...)

	return nil
}

// GetHeap is a function to return a pointer to all tuples in the heap
func (heap *Heap) GetHeap() []*Tuple {
	return heap.Tuples
}
