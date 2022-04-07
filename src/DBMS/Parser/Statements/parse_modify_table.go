package Statements

import (
	"GoDBMS/ParserStructs"
    "strings"
	"errors"
)

func ParseModifyTable(query string) (*ParserStructs.ModifyTableStatement, error) {
	
	query = strings.Trim(query, " ;")

	querySplit := strings.Split(query, " ")

	tableName := ""
	columnName := ""
	dataType := ""
	

	//check if statement has correct number of arguments
	if len(querySplit) != 6 {
		return nil, errors.New("Invalid alter table statement")
	}

	
	//check if valid syntax/ which type of alter table statement it is
	if querySplit[3] == "add" {
		tableName = querySplit[2]
		columnName = querySplit[4]
		dataType = querySplit[5]

		if dataType != "int" && dataType != "string" {
			return nil, errors.New("Invalid column datatype")
		}

	} else if querySplit[3] == "drop" && querySplit[4] == "column"{
		tableName = querySplit[2]
		columnName = querySplit[5]
	} else {
		return nil, errors.New("Invalid alter table statement")
	}

	return &ParserStructs.ModifyTableStatement{tableName, columnName, dataType}, nil
}