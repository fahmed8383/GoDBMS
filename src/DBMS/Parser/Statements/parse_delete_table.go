package Statements

import (
	"GoDBMS/ParserStructs"
    "strings"
	"errors"
)

func ParseDeleteTable(query string) (*ParserStructs.DeleteTableStatement, error) {
	
	querySplit := strings.Split(query, " ")
	name := querySplit[2]
	
	return &ParserStructs.DeleteTableStatement{name}, nil

	

