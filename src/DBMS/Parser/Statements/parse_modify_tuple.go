package Statements

import (
	"GoDBMS/ParserStructs"
	"errors"
	"strings"
)

// ParseModifyTuple is a function that is used to parse a query and return a
// ModifyTupleStatement struct.
func ParseModifyTuple(query string) (*ParserStructs.ModifyTupleStatement, error) {
	
	// Replace set with ? so we can split over the ? delimeter
	queryReplaceSet := strings.ReplaceAll(query, "set", "?")
	querySplitSet := strings.Split(queryReplaceSet, "?")

	// Check to make sure we have a set clause
	if len(querySplitSet) != 2 {
		return nil, errors.New("Invalid update statement syntax")
	}

	// Check to make sure we have a table name
	queryTableTrim := strings.Trim(querySplitSet[0], " ,;")
	queryTableSplit := strings.Split(queryTableTrim, " ")
	if len(queryTableSplit) != 2 || queryTableSplit[1] == "" {
		return nil, errors.New("Update statement is missing  the table name")
	}

	// Save table name and initialize set and where arrays
	tableName := strings.Trim(queryTableSplit[1], " ,;")
	setArray := [][2]string{}
	whereArray := []string{}

	// Replace where with ? so we can split over the ? delimeter
	queryReplaceWhere := strings.ReplaceAll(querySplitSet[1], "where", "?")
	querySplitWhere := strings.Split(queryReplaceWhere, "?")

	// Split over all set parameters
	querySetParams := strings.Split(querySplitWhere[0], ",")
	for _, param := range querySetParams {

		// Check to make sure set parameters have the appropriate syntax
		paramVal := strings.Split(param, "=")
		if len(paramVal) != 2 {
			return nil, errors.New("Update statement has invalid set parameter syntax")
		}

		paramArray := [2]string{strings.Trim(paramVal[0], " ,;"), strings.Trim(paramVal[1], " ,;")}
		setArray = append(setArray, paramArray)
	}

	// Check to make sure that we have atleast one set parameter
	if len(setArray) == 0 {
		return nil, errors.New("Update statement is missing set parameters")
	}

	// Make sure where clause actually exists
	if len(querySplitWhere) > 1 {

		whereTrim := strings.Trim(querySplitWhere[1], " ,;")
		
		// If where parameter does not have space between operator, return error
		// asking them to put the appropriate spaces.
		if strings.Count(whereTrim, " ") != 2 {
			return nil, errors.New("Please put spaces before and after the comparator in the where clause")
		}

		// Make sure where parameter has two values and a comparator
		whereSplit := strings.Split(whereTrim, " ")
		if len(whereSplit) != 3 {
			return nil, errors.New("Invalid where parameter")
		}

		valueOne := strings.Trim(whereSplit[0], " ,;")
		comparator := strings.Trim(whereSplit[1], " ,;")
		valueTwo := strings.Trim(whereSplit[2], " ,;")

		// Make sure we have a valid comparator
		if comparator != ">" && comparator != "<" && comparator != "=" && comparator != ">=" && comparator != "<=" {
			return nil, errors.New("Invalid comparator")
		}

		// Add values and comparator to the array
		whereArray = append(whereArray, valueOne, comparator, valueTwo)
	}

	return &ParserStructs.ModifyTupleStatement{tableName, setArray, whereArray}, nil
}