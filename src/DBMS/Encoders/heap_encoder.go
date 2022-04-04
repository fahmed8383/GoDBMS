package Encoders

import (
	"GoDBMS/Storage"
	"bytes"
	"encoding/gob"
)

// EncodeHeap is a function to serialize the catalog data structure into
// bytes that can be written to a data file.
func EncodeHeap(name string, heapPointer *Storage.Heap) error {

	// Encode the heap pointer into a bytes buffer.
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(heapPointer)

	// Write the bytes from the buffer into a byte data file and log any errors
	// that occur during the write.
	err := WriteByteFile(name, buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// DecodeHeap is a function to deserialize the heap data structure from
// a byte data file on the disk.
func DecodeHeap(name string) (*Storage.Heap, error) {

	// Check if the heap data file exists.
	if FileExists(name) {

		// If it does exist, read the bytes from the file and log any errors
		// that occur during the read.
		heapBytes, err := ReadByteFile(name)
		if err != nil {
			return nil, err
		}

		// Write the bytes into a byte buffer.
		buffer := bytes.NewBuffer(heapBytes)

		// Initialize the heap datastructure.
		var heap Storage.Heap

		// Decode the byte buffer into the heap data structure.
		dec := gob.NewDecoder(buffer)
		dec.Decode(&heap)

		return &heap, nil

	} else {

		// If the heap does not already exist, initliaze an empty heap
		// in memory.
		return Storage.InitializeHeap(), nil
	}

	return nil, nil
}
