package Statements

import (
	"GoDBMS/ParserStructs"
	"strings"
	"errors"
)

// ParseSelect is a function that is used to parse the select query and return
// a pointer to the SelectStatement struct and any errors.
func ParseSelect(query string) (*ParserStructs.SelectStatement, error) {

	// Replace select with ? so we can split over the ? delimeter
	queryReplaceSelect := strings.ReplaceAll(query, "select", "?")
	querySplitSelect := strings.Split(queryReplaceSelect, "?")
	if len(querySplitSelect) != 2 {
		return nil, errors.New("Missing select in select statement")
	}

	// Replace from with ? so we can split over the ? delimeter
	queryReplaceFrom := strings.ReplaceAll(querySplitSelect[1], "from", "?")
	querySplitFrom := strings.Split(queryReplaceFrom, "?")
	if len(querySplitFrom) != 2 {
		return nil, errors.New("Missing from in select statement")
	}

	// Make sure select clause has atleast one parameter
	selectSplit := strings.Split(querySplitFrom[0], ",")
	if len(selectSplit) == 1 && selectSplit[0] == " " {
		return nil, errors.New("Missing select parameters")
	}

	selectArray := []string{}
	fromString := ""
	whereArray := []string{}

	// Add all select parameters to an array
	for i := 0; i < len(selectSplit); i++ {
		value := strings.Trim(selectSplit[i], " ,;")
		selectArray = append(selectArray, value)
	}

	// Replace where with ? so we can split over the ? delimeter
	queryReplaceWhere := strings.ReplaceAll(querySplitFrom[1], "where", "?")
	querySplitWhere := strings.Split(queryReplaceWhere, "?")

	// Make sure from clause has one nonempty parameter
	fromTrim := strings.Trim(querySplitWhere[0], " ,;")
	fromSplit := strings.Split(fromTrim, " ")
	if len(fromSplit) != 1 || fromSplit[0] == " " || fromSplit[0] == "" {
		return nil, errors.New("Invalid from parameter")
	}

	fromString = fromSplit[0]

	// Make sure where clause actually exists
	if len(querySplitWhere) > 1 {

		// Make sure where parameter has two values and a comparator
		whereTrim := strings.Trim(querySplitWhere[1], " ,;")
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
	
	return &ParserStructs.SelectStatement{selectArray, fromString, whereArray}, nil
}
