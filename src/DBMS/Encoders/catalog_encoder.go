package Encoders

import (
	"GoDBMS/Storage"
	"log"
	"encoding/gob"
	"bytes"
)

// EncodeCatalog is a function to serialize the catalog data structure into
// bytes that can be written to a data file.
func EncodeCatalog() {

	// Get a pointer to the current catalog in memory.
	catalogPointer := Storage.GetTablesMap()
	
	// Encode the catalog pointer into a bytes buffer.
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(catalogPointer)

	// Write the bytes from the buffer into a byte data file and log any errors
	// that occur during the write.
	err := WriteByteFile("catalog", buffer.Bytes())
	if err != nil {
		log.Println(err)
	}
}

// DecodeCatalog is a function to deserialize the catalog data structure from
// a byte data file on the disk.
func DecodeCatalog() {

	// Check if the catalog data file exists.
	if FileExists("catalog") {

		// If it does exist, read the bytes from the file and log any errors
		// that occur during the read.
		catalogBytes, err := ReadByteFile("catalog")
		if err != nil {
			log.Println(err)
		}

		// Write the bytes into a byte buffer.
		buffer := bytes.NewBuffer(catalogBytes)
		
		// Initialize the catalog datastructure.
		var catalog map[string]*Storage.TableSchema

		// Decode the byte buffer into the catalog data structure.
		dec := gob.NewDecoder(buffer)
		dec.Decode(&catalog)

		// Load the catalog into memory.
		Storage.LoadTablesMap(&catalog)
		
	} else {
		
		// If the catalog does not already exist, initliaze an empty catalog
		// in memory.
		Storage.InitializeTables()
	}
}