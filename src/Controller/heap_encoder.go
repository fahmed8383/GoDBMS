import (
	"GoDBMS/Database"
	"log"
	"encoding/gob"
	"bytes"
)

// EncodeHeap is a function to serialize the catalog data structure into
// bytes that can be written to a data file.
func EncodeHeap() {

	// Get a pointer to the current heap in memory.
	heapPointer := Database.GetHeap()
	
	// Encode the heap pointer into a bytes buffer.
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(heapPointer)

	// Write the bytes from the buffer into a byte data file and log any errors
	// that occur during the write.
	err := WriteByteFile("heap", buffer.Bytes())
	if err != nil {
		log.Println(err)
	}
}