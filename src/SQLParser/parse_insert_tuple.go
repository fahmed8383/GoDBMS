package SQLParser

import (
	"errors"
	"strings"
)

// InsertTupleStatement holds the insert tuple query info received from the user
type InsertTupleStatement struct {
	// TableName is a string representing name of the table
	TableName string
	// Columns is an array of tableColumn representing the columns
	Columns []tableColumn
}

// tableColumn holds the insert tuple query column info received from the user
type tableColumn struct {
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
	}

	// Extracts the table name and coverts it to lowercase
	name := strings.Split(querySplit[0], " ")[2]
	name = strings.ToLower(name)

	// Extracts the column names and trims them
	columnsSplit := strings.Split(querySplit[1], " ")
	// Removes irrelevant words from the column names
	getColumns := columnsSplit[0:len(querySplit)]
	for i, v := range getColumns {
		getColumns[i] = strings.Trim(v, " ,)")
	}

	// Extracts the values for each corresponding column name and trims it
	getValues := strings.Split(querySplit[2], " ")
	for i, v := range getValues {
		getValues[i] = strings.Trim(v, " ,);")
	}

	// Initializes the array of tableColumns
	columns := []tableColumn{}
	for i := range getColumns {
		columnStruct := tableColumn{getColumns[i], getValues[i]}
		columns = append(columns, columnStruct)
	}

	// Returns a new InsertTupleStatement struct pointer based on the query
	return &InsertTupleStatement{name, columns}, nil
}
