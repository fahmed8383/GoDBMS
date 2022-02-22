package SQLParser

import (
    "strings"
	"errors"
)

// CreateTableStatement holds the create table query info received from the
// user.
type CreateTableStatement struct {
	// Name is a string representing name of the table.
	Name string
	// PrimaryKeyIndex is an int representing the index of the primary key 
	// column.
	PrimaryKeyIndex int
	// PrimaryKeyIndex is an array of createTableColumn representing the 
	// columns.
	Columns []createTableColumn
}

// createTableColumn holds the create table query column info received from the
// user.
type createTableColumn struct {
	// Name is a string representing name of the column.
	Name string
	// Name is a string representing datatype of the column.
	Datatype string
	// NotNull is a boolean representing whether the entries into the column
	// can be null or not.
	NotNull bool
}

// parseTableQuery is a function that is used to parse the create table query
// input into the CreateTableStatement struct.
func parseTableQuery(query string) (*CreateTableStatement, error) {

	// Split the user input at open bracket to seperate the table name and
	// column info
	bracketSplit := strings.Split(query, "(")

	// Get the table name and convert it to lowercase
	name := strings.Split(bracketSplit[0], " ")[2]
	name = strings.ToLower(name)

	// Split the column info at coma to seperate all the columns
	columnsSplit := strings.Split(bracketSplit[1], ",")

	primaryKeyIndex := -1
	columns := []createTableColumn{}

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
			return nil, errors.New("Create table statement has too many"+
									" parameters")
		}
		if len(columnData) < 2 {
			return nil, errors.New("Create table statement must have a name"+ 
									" and a type for each column")
		}

		// Get column name and type and set them to lowercase
		columnName := strings.ToLower(columnData[0])
		columnType := strings.ToLower(columnData[1])
		notNull := false

		// Check to make sure that the column type is a valid datatype
		if columnType != "string" && columnType != "int" {
			return nil, errors.New("Create table statement must have an int or"+
									" string datatype")
		}

		// Check to see if an optional parameter is included for the column
		if len(columnData) > 2 {
			// Convert the optional parameter to lower case
			optParam := strings.ToLower(columnData[2])

			// Check to see if it is a valid optional parameter
			switch optParam {

			case "not null":
				notNull = true

			case "not null primary key", "primary key not null", "primary key":

				// Ensure that the user input does not have more than one 
				// primary key
				if primaryKeyIndex != -1 {
					return nil, errors.New("Create table statement cannot have"+ 
											" more than one primary key")
				}
				primaryKeyIndex = i
				notNull = true

			default:
				return nil, errors.New("Create table statement has an invalid"+
										" parameter")
			}
		}
		
		// Create the column struct from the user input and add it to the 
		// columns array
		columnStruct := createTableColumn{columnName, columnType, notNull}
		columns = append(columns, columnStruct)
	}

	// Ensure that the user input has a primary key
	if primaryKeyIndex == -1 {
		return nil, errors.New("Create table statement missing a primary key")
	}

	// Create and return the CreateTableStatement struct pointer from the user
	// input
	return &CreateTableStatement{name, primaryKeyIndex, columns}, nil
}