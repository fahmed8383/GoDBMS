package Statements

import (
	"GoDBMS/ParserStructs"
    "strings"
	"errors"
)

// ParseDeleteTable is a function that is used to parse the delete table query
// and return a pointer to the DeleteTableStatment struct and any errors.
func ParseDeleteTable(query string) (*ParserStructs.DeleteTableStatement, error) {
	
	query = strings.Trim(query, " ;")

	querySplit := strings.Split(query, " ")

	if len(querySplit) != 3 {
		return nil, errors.New("Delete table statement does not contain a table name")
	}

	name := querySplit[2]

	return &ParserStructs.DeleteTableStatement{name}, nil
}