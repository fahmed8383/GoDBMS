package ProcessSQLStatements

import (
	"GoDBMS/ParserStructs"
	"GoDBMS/Storage"
	"GoDBMS/Encoders"
	"strconv"
	"errors"
)

// ProcessSelect is a function that handles the searching of a table from
// the database. It takes an InsertTupleStatement pointer as an argument a
// a string with the search results and any errors
func ProcessSelect(selectStatement *ParserStructs.SelectStatement) (string, error) {

	// Check if where is present in select statement
	whereExists := false
	whereIndex := -1
	if len(selectStatement.Where) != 0 {
		whereExists = true
	}

	// Check if star operator is used in selection
	starOperator := false
	for _, column := range selectStatement.Select {
		if column == "*" {
			starOperator = true
		}
	}

	tableName := selectStatement.From

	// Check if table exists, if not, return error
	if !Storage.TableExists(tableName) {
		return "", errors.New("Table " + tableName + " from select statement does not exist")
	}

	// Get the table corresponding to the from clause
	table := Storage.GetTable(tableName)

	// If whereExists the ensure the column specified in where is valid
	if whereExists {
		index, exists := table.ColumnIndex[selectStatement.Where[0]]
		whereIndex = index
		if !exists {
			return "", errors.New("Column " + selectStatement.Where[0] + " not found in table")
		}
		if table.Columns[index].Datatype == "string" && selectStatement.Where[1] != "=" {
			return "", errors.New("Operator " + selectStatement.Where[1]+ " is not supported for strings")
		}
	}

	// If star operator does not exist, ensure that all columns in the selection
	// exist
	if !starOperator {
		for _, column := range selectStatement.Select {
			_, exists := table.ColumnIndex[column]
			if !exists {
				return "", errors.New("Column " + column + " not found in table")
			}
		}
	}

	// Save an array of valid tuple values that meet the where statement
	// requirements
	validTuples := [][]interface{}{}

	// Get the heap for the table. This contains all the tuple values
	heap, err := Encoders.DecodeHeap(tableName)
	if err != nil {
		return "", err
	}

	for _, tuple := range heap.GetHeap() {
		// If where does not exist, then all values are valid
		if !whereExists {
			validTuples = append(validTuples, tuple.Values)

		// Otherwise add values that meet the where specification to the valid
		// array
		} else {
			if table.Columns[whereIndex].Datatype == "string" {
				stringVal := tuple.Values[whereIndex].(string)
				if stringVal == selectStatement.Where[2] {
					validTuples = append(validTuples, tuple.Values)
				}
			} else {
				intVal := tuple.Values[whereIndex].(int)
				compareVal, err := strconv.Atoi(selectStatement.Where[2])
				if err != nil {
					return "", errors.New("Where comparison parameter cannot be converted to an int")
				}

				if selectStatement.Where[1] == "=" && intVal == compareVal {
					validTuples = append(validTuples, tuple.Values)
				} else if selectStatement.Where[1] == ">" && intVal > compareVal {
					validTuples = append(validTuples, tuple.Values)
				} else if selectStatement.Where[1] == "<" && intVal < compareVal {
					validTuples = append(validTuples, tuple.Values)
				} else if selectStatement.Where[1] == ">=" && intVal >= compareVal {
					validTuples = append(validTuples, tuple.Values)
				} else if selectStatement.Where[1] == "<=" && intVal <= compareVal {
					validTuples = append(validTuples, tuple.Values)
				}
			}
		}
	}

	// Create a string of valid column names
	colNames := "\n"
	if starOperator {
		for i, tableCol := range table.Columns {
			if i == 0 {
				colNames += " " + tableCol.Name + " "
			} else {
				colNames += "| " + tableCol.Name + " "
			}
		}
	} else {
		for i, colName := range selectStatement.Select {
			if i == 0 {
				colNames += " " + colName + " "
			} else {
				colNames += "| " + colName + " "
			}
		}
	}

	// Add dashes after column names
	dashes := ""
	for j := 0; j < len(colNames); j++ {
		dashes += "-"
	}

	outputString := "\n" + colNames + "\n" + dashes + "\n"

	// Loop over all valid tuples and add them to the output string
	for _, tuple := range validTuples {
		tupleString := ""
		if starOperator {
			for i, val := range tuple {
				stringVal := getStringVal(val, table.Columns[i].Datatype)
				if i == 0 {
					tupleString += " " + stringVal + " "
				} else {
					tupleString += "| " + stringVal + " "
				}
			}
		} else {
			for i, colName := range selectStatement.Select {
				colIndex := table.ColumnIndex[colName]
				stringVal := getStringVal(tuple[colIndex], table.Columns[colIndex].Datatype)
				if i == 0 {
					tupleString += " " + stringVal + " "
				} else {
					tupleString += "| " + stringVal + " "
				}
			}
		}

		outputString += tupleString + "\n"
	}

	return outputString, nil
}

// getStringVal takes in an interface value and the datatype and returns the corresponding
// string value
func getStringVal(val interface{}, dataType string) string {
	if dataType == "string" {
		return val.(string)
	} else {
		intVal := val.(int)
		return strconv.Itoa(intVal)
	}
}