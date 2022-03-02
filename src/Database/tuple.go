package Database

import (
	"strconv"
)

// A struct that describes the schema of a Tuple
type tuple struct {
	Values []interface{}
}

// Constructor function to initialize a Tuple and return a pointer to the tuple
func CreateTuple(table *TableSchema, insertQuery *InsertTupleStatement) (*tuple, error) {

	// Creates a new array with interface{} type to accept multiple data types
	values := []interface{}{}

	// Iterates through the columns of the values to check data types
	for i, column := range insertQuery.Columns {

		// If the data type is int, then the value will be converted to int and added to the array
		if table.Columns[i].Datatype == "int" {
			val, err := strconv.Atoi(column.Value)
			// If there is an error, then the value is not an int and the function returns an error
			if err != nil {
				return nil, err
			}
			values = append(values, val)
			// If the data type is string, then the value will be added to the array
		} else if table.Columns[i].Datatype == "string" {
			values = append(values, column.Value)
		}
	}

	// Creates a new tuple and returns a pointer to the tuple
	return &tuple{values}, nil
}
