package SQLParser

import (
	"errors"
	"strings"
)

// InsertTupleStatement holds the insert tuple query info received from the user
type InsertTupleStatement struct {
	// TableName is a string representing name of the table
	TableName string
	// Columns is an array of insertTupleColumn representing the columns
	Columns []insertTupleColumn
}

// insertTupleColumn holds the insert tuple query column info received from the user
type insertTupleColumn struct {
	// Name is a string representing name of the column
	Name string
	// Value is a string representing the value of the column
	Value string
}

// ParseInsertTupleQuery is a function that is used to parse a query
// and return a InsertTupleStatement struct.
func parseInsertTupleQuery(query string) (*InsertTupleStatement, error) {
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

	// Extracts the values for each corresponding column name and trims it
	getValues := strings.Split(querySplit[2], ",")
	for i, v := range getValues {
		getValues[i] = strings.Trim(v, " );")
	}

	// Initializes the array of insertTupleColumns
	columns := []insertTupleColumn{}

	// Checks if the number of columns and values are the same
	if len(getColumns) != len(getValues) {
		return nil, errors.New("Insert tuple statement has different number of columns and values")
	}

	// Goes through each column name and value, creates a insertTupleColumn struct
	// and then adds them into the array of insertTupleColumns
	for i := range getValues {
		columnStruct := insertTupleColumn{getColumns[i], getValues[i]}
		columns = append(columns, columnStruct)
	}

	// Returns a new InsertTupleStatement struct pointer based on the query
	return &InsertTupleStatement{name, columns}, nil
}
