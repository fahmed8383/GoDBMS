package SQLParser

import (
	"GoDBMS/Database"
	"errors"
	"strings"
)

// ParseInsertTupleQuery is a function that is used to parse a query
// and return a InsertTupleStatement struct.
func parseInsertTupleQuery(query string) (*Database.InsertTupleStatement, error) {
	// Divides the query into three parts: the table name, the column names, and the
	// values
	querySplit := strings.Split(query, "(")

	// Checks if the query has the correct number of arguments
	if len(querySplit) > 3 {
		return nil, errors.New("Insert tuple statement has too many arguments")
	} else if len(querySplit) < 3 {
		return nil, errors.New("Insert tuple statement has too few arguments")
	}

	// Extracts the table name and coverts it to lowercase
	nameSplit := strings.Split(querySplit[0], " ")
	// Checks if the query has the correct format for the table name
	if len(nameSplit) != 4 {
		return nil, errors.New("Insert tuple statement has invalid table name format")
	}
	name := nameSplit[2]

	// Extracts the column section and trims them
	columnSplit := strings.Split(querySplit[1], ")")

	// Checks if the query has the correct syntax for insert statement
	if len(columnSplit) < 2 {
		return nil, errors.New("Insert tuple statement has invalid syntax")
	}

	if strings.Trim(columnSplit[1], " ") != "values" {
		return nil, errors.New("Insert tuple statement is missing values keyword")
	}

	// Removes irrelevant words from the column names
	// e.g. "[column1, column2),  VALUES]" -> "[column1, column2)]"
	columnTrim := columnSplit[0]
	// Gets the column names by splitting with commas
	getColumns := strings.Split(columnTrim, ",")
	for i, col := range getColumns {
		getColumns[i] = strings.Trim(col, " ")
	}

	// Extracts the values for each corresponding column name and trims it
	getValues := strings.Split(querySplit[2], ",")
	for i := range getValues {
		getValues[i] = strings.Trim(getValues[i], " );")
	}

	// Initializes the array of InsertTupleColumns
	columns := []Database.InsertTupleColumn{}

	// Checks if the number of columns and values are the same
	if len(getColumns) != len(getValues) {
		return nil, errors.New("Insert tuple statement has different number of columns and values")
	}

	// Goes through each column name and value, creates a InsertTupleColumn struct
	// and then adds them into the array of InsertTupleColumns
	for i := range getValues {
		columnStruct := Database.InsertTupleColumn{getColumns[i], getValues[i]}
		columns = append(columns, columnStruct)
	}

	// Returns a new InsertTupleStatement struct pointer based on the query
	return &Database.InsertTupleStatement{name, columns}, nil
}
