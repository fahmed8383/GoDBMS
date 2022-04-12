package ProcessSQLStatements

import (
	"GoDBMS/ParserStructs"
	"GoDBMS/Storage"
	"GoDBMS/StorageLock"
	"GoDBMS/Encoders"
	"strconv"
	"errors"
)

// ProcessDeleteTuple is a function that takes in a pointer to the
// DeleteTupleStatement struct and deletes that tuple from its appropriate heap
// file. It returns any errors that occur.
func ProcessDeleteTuple(deleteTuple *ParserStructs.DeleteTupleStatement) error {

	// Check if where is present in delete statement
	whereExists := false
	whereIndex := -1
	if len(deleteTuple.Where) != 0 {
		whereExists = true
	}

	tableName := deleteTuple.From

	// Check if table exists, if not, return error
	if !Storage.TableExists(tableName) {
		return errors.New("Table " + tableName + " from delete statement does not exist")
	}

	// Get the table corresponding to the from clause
	table := Storage.GetTable(tableName)

	StorageLock.AcquireTableLock(tableName)

	// If where does not exist, then all tuples need to be deleted
	if !whereExists {
		heap := Storage.InitializeHeap()

		// Encode and save the new empty heap
		err := Encoders.EncodeHeap(tableName, heap)
		if err != nil {
			StorageLock.ReleaseTableLock(tableName)
			return err
		}
		StorageLock.ReleaseTableLock(tableName)
		return nil
	}

	// If whereExists the ensure the column specified in where is valid
	index, exists := table.ColumnIndex[deleteTuple.Where[0]]
	whereIndex = index
	if !exists {
		StorageLock.ReleaseTableLock(tableName)
		return errors.New("Column " + deleteTuple.Where[0] + " not found in table")
	}
	if table.Columns[index].Datatype == "string" && deleteTuple.Where[1] != "=" {
		StorageLock.ReleaseTableLock(tableName)
		return errors.New("Operator " + deleteTuple.Where[1]+ " is not supported for strings")
	}

	// Get the heap for the table. This contains all the tuple values
	heap, err := Encoders.DecodeHeap(tableName)
	if err != nil {
		StorageLock.ReleaseTableLock(tableName)
		return err
	}

	// Loop over all tuples
	for _, tuple := range heap.GetHeap() {
		keyInd := table.PrimaryKeyIndex
		keyVal := tuple.Values[keyInd]

		// Remove the tuples that fit the where specification
		if table.Columns[whereIndex].Datatype == "string" {
			stringVal := tuple.Values[whereIndex].(string)
			if stringVal == deleteTuple.Where[2] {
				heap.DeleteTuple(keyVal, keyInd)
			}
		} else {
			intVal := tuple.Values[whereIndex].(int)
			compareVal, err := strconv.Atoi(deleteTuple.Where[2])
			if err != nil {
				StorageLock.ReleaseTableLock(tableName)
				return errors.New("Where comparison parameter cannot be converted to an int")
			}

			if deleteTuple.Where[1] == "=" && intVal == compareVal {
				heap.DeleteTuple(keyVal, keyInd)
			} else if deleteTuple.Where[1] == ">" && intVal > compareVal {
				heap.DeleteTuple(keyVal, keyInd)
			} else if deleteTuple.Where[1] == "<" && intVal < compareVal {
				heap.DeleteTuple(keyVal, keyInd)
			} else if deleteTuple.Where[1] == ">=" && intVal >= compareVal {
				heap.DeleteTuple(keyVal, keyInd)
			} else if deleteTuple.Where[1] == "<=" && intVal <= compareVal {
				heap.DeleteTuple(keyVal, keyInd)
			}
		}
	}

	// Encode and save the new modified heap
	err = Encoders.EncodeHeap(tableName, heap)
	if err != nil {
		StorageLock.ReleaseTableLock(tableName)
		return err
	}

	StorageLock.ReleaseTableLock(tableName)
	return nil
}