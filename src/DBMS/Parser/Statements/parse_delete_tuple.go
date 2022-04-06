package Statements

import (
	"GoDBMS/ParserStructs"
    "strings"
	"errors"
)

// ParseDeleteTuple is a function that is used to parse the delete query and
// return a pointer to the DeleteTableStatment struct and any errors.
func ParseDeleteTuple(query string) (*ParserStructs.DeleteTupleStatement, error) {

	// Replace where with ? so we can split over the ? delimeter
	queryReplaceWhere := strings.ReplaceAll(query, "where", "?")
	querySplitWhere := strings.Split(queryReplaceWhere, "?")

	// Make sure we have correct string before the where clause
	queryTrim := strings.Trim(querySplitWhere[0], " ,;")
	querySplit := strings.Split(queryTrim, " ")
	if len(querySplit) != 3 {
		return nil, errors.New("Invalid delete tuple statement")
	}
	if querySplit[0] != "delete" || querySplit[1] != "from" {
		return nil, errors.New("Invalid delete tuple syntax")
	}

	fromString := strings.Trim(querySplit[2], " ,;")
	whereArray := []string{}

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

	return &ParserStructs.DeleteTupleStatement{fromString, whereArray}, nil
}