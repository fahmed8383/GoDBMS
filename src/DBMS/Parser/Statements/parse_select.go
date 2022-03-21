package Statements

import (
	"GoDBMS/ParserStructs"
	"strings"
	"errors"
)

func ParseSelect(query string) (*ParserStructs.SelectStatement, error) {
	querySplit := strings.Split(query, " ")
	if len(querySplit) != 4 {
		return nil, errors.New("Invalid number of parameters for select statement")
	}
	if querySplit[1] != "*" || querySplit[2] != "from" {
		return nil, errors.New("Invalid syntax for select statement")
	}

	name := querySplit[3]
	
	return &ParserStructs.SelectStatement{name}, nil
}

