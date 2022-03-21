package Statements

import (
	"GoDBMS/ParserStructs"
    "strings"
	"errors"
)

// ParseCreateTable is a function that is used to parse the create table query
// input into the CreateTableStatement struct.
func ParseCreateTable(query string) (*ParserStructs.CreateTableStatement, error) {

	// Split the user input at open bracket to seperate the table name and
	// column info
	bracketSplit := strings.Split(query, "(")

	// Remove all spaces on the ends of the table name info string and split it
	// at all remaining spaces in between.
	nameTrim := strings.Trim(bracketSplit[0], " ")
	nameSplit := strings.Split(nameTrim, " ")

	// If the array of strings that we get by splitting it at space is not equal
	// to three, this means that either we are missing a table name or the
	// table name is multple words.
	if len(nameSplit) != 3 {
		return nil, errors.New("Create table query has an invalid table name")
	}

	// Get the table name and convert it to lowercase
	name := nameSplit[2]

	// Split the column info at coma to seperate all the columns
	columnsSplit := strings.Split(bracketSplit[1], ",")

	primaryKeyIndex := -1
	columns := []ParserStructs.CreateTableColumn{}

	// Loop through all the columns
	for i, columnString := range columnsSplit {

		// Clean up the string data for each column and seperate each piece of
		// info into a maximum of three strings. The first two strings will
		// be the name and type of the column, while the last one would be the
		// optional parameter.
		columnString = strings.Trim(columnString, "  ,);")
		columnData := strings.SplitN(columnString, " ", 3)
		
		// Check for syntax errors in user input
		if len(columnData) == 0 {
			return nil, errors.New("Create table statement has an extra comma")
		}
		if len(columnData) > 3 {
			return nil, errors.New("Create table statement has too many parameters")
		}
		if len(columnData) < 2 {
			return nil, errors.New("Create table statement must have a name and a type for each column")
		}

		// Get column name and type and set them to lowercase
		columnName := columnData[0]
		columnType := columnData[1]
		notNull := false

		// Check to make sure that the column type is a valid datatype
		if columnType != "string" && columnType != "int" {
			return nil, errors.New("Create table statement must have an int or string datatype")
		}

		// Check to see if an optional parameter is included for the column
		if len(columnData) > 2 {
			// Convert the optional parameter to lower case
			optParam := columnData[2]

			// Check to see if it is a valid optional parameter
			switch optParam {

			case "not null":
				notNull = true

			case "not null primary key", "primary key not null", "primary key":

				// Ensure that the user input does not have more than one 
				// primary key
				if primaryKeyIndex != -1 {
					return nil, errors.New("Create table statement cannot have more than one primary key")
				}
				primaryKeyIndex = i
				notNull = true

			default:
				return nil, errors.New("Create table statement has an invalid parameter")
			}
		}
		
		// Create the column struct from the user input and add it to the 
		// columns array
		columnStruct := ParserStructs.CreateTableColumn{columnName, columnType, notNull}
		columns = append(columns, columnStruct)
	}

	// Ensure that the user input has a primary key
	if primaryKeyIndex == -1 {
		return nil, errors.New("Create table statement missing a primary key")
	}

	// Create and return the CreateTableStatement struct pointer from the user
	// input
	return &ParserStructs.CreateTableStatement{name, primaryKeyIndex, columns}, nil
}