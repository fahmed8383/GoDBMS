package ProcessSQLStatements

import (
	"GoDBMS/ParserStructs"
	"GoDBMS/Storage"
	"GoDBMS/StorageLock"
	"GoDBMS/Encoders"
	"strconv"
	"errors"
)

// ProcessModifyTuple is a function that handles the modification/update of a
// tuple in the heap. It takes an ModifyTupleStatement pointer as an argument
// and returns any errors if they exist.
func ProcessModifyTuple(modifyTuple *ParserStructs.ModifyTupleStatement) error {

	// Check if where is present in modify tuple statement
	whereExists := false
	whereIndex := -1
	if len(modifyTuple.Where) != 0 {
		whereExists = true
	}

	tableName := modifyTuple.TableName

	// Check if table exists, if not, return error
	if !Storage.TableExists(tableName) {
		return errors.New("Table " + tableName + " from update statement does not exist")
	}

	// Get the table from the catalog
	table := Storage.GetTable(tableName)

	StorageLock.AcquireTableLock(tableName)

	// Check to make sure all columns in the set param exist
	for _, setParam := range modifyTuple.Set {
		column := setParam[0]
		_, exists := table.ColumnIndex[column]
		if !exists {
			StorageLock.ReleaseTableLock(tableName)
			return errors.New("Column " + column + " not found in table")
		}
	}

	// If whereExists the ensure the column specified in where is valid
	if whereExists {
		index, exists := table.ColumnIndex[modifyTuple.Where[0]]
		whereIndex = index
		if !exists {
			StorageLock.ReleaseTableLock(tableName)
			return errors.New("Column " + modifyTuple.Where[0] + " not found in table")
		}
		if table.Columns[index].Datatype == "string" && modifyTuple.Where[1] != "=" {
			StorageLock.ReleaseTableLock(tableName)
			return errors.New("Operator " + modifyTuple.Where[1]+ " is not supported for strings")
		}
	}

	// Get the heap for the table. This contains all the tuple values
	heap, err := Encoders.DecodeHeap(tableName)
	if err != nil {
		StorageLock.ReleaseTableLock(tableName)
		return err
	}

	// Loop over all tuples
	for _, tuple := range heap.GetHeap() {

		// If where does not exist then modify the columns for all tuples
		if !whereExists {
			for _, setParam := range modifyTuple.Set {
				colIndex := table.ColumnIndex[setParam[0]]
				err := modifyTupleValue(tuple, colIndex, table.Columns[colIndex].Datatype, setParam[1])
				if err != nil {
					StorageLock.ReleaseTableLock(tableName)
					return errors.New("Unable to convert set parameter value to int")
				}
			}

		// otherwise, modify the columns for only the tuples that match the where parameters
		} else {
			matches, err := matchesWhere(tuple.Values[whereIndex], table.Columns[whereIndex].Datatype, modifyTuple.Where[1], modifyTuple.Where[2])
			if err != nil {
				StorageLock.ReleaseTableLock(tableName)
				return errors.New("Unable to convert where parameter value to int")
			}

			if matches {
				for _, setParam := range modifyTuple.Set {
					colIndex := table.ColumnIndex[setParam[0]]
					err := modifyTupleValue(tuple, colIndex, table.Columns[colIndex].Datatype, setParam[1])
					if err != nil {
						StorageLock.ReleaseTableLock(tableName)
						return errors.New("Unable to convert set parameter value to int")
					}
				}
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

// Takes in a tuple, a column index, datatype, and a new val and uses it to
// modify the tuple
func modifyTupleValue(tuple *Storage.Tuple, i int, dataType string, val string) error {
	if dataType == "string" {
		tuple.Values[i] = val
	} else {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		tuple.Values[i] = intVal
	}

	return nil
}

// Takes in a tuple value, a datatype, and information about the hwer statement
// check if the current value is included in the where clause
func matchesWhere(tupleVal interface{}, dataType string, operator string, whereVal string) (bool, error) {

	if dataType == "string" {
		stringVal := tupleVal.(string)
		if stringVal == whereVal {
			return true, nil
		}
	} else {
		intVal := tupleVal.(int)
		compareVal, err := strconv.Atoi(whereVal)
		if err != nil {
			return false, err
		}

		if operator == "=" && intVal == compareVal {
			return true, nil
		} else if operator == ">" && intVal > compareVal {
			return true, nil
		} else if operator == "<" && intVal < compareVal {
			return true, nil
		} else if operator == ">=" && intVal >= compareVal {
			return true, nil
		} else if operator == "<=" && intVal <= compareVal {
			return true, nil
		}
	}

	return false, nil
}