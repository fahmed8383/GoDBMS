package Statements

import (
	"GoDBMS/ParserStructs"
    "strings"
	"errors"
)

func ParseDeleteTable(query string) (*ParserStructs.DeleteTableStatement, error) {
	
	query = strings.Trim(query, " ;")

	querySplit := strings.Split(query, " ")

	if len(querySplit) != 3 {
		return nil, errors.New("Delete table statement does not contain a table name")
	}

	name := querySplit[2]

	return &ParserStructs.DeleteTableStatement{name}, nil
}