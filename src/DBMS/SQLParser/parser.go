package SQLParser

import (
	"GoDBMS/Controller"
	"fmt"
	"strings"
)

func ParseInput(query string) (string) {

	querySplit := strings.Split(query, " ")

	//if the user input is a 'create table' query call parseCreateTableQuery
	if querySplit[0] == "create" && querySplit[1] == "table" {
		output, err := parseCreateTableQuery(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		err = Controller.ProcessCreateTable(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return "Table created successfully"

	} else if querySplit[0] == "insert" && querySplit[1] == "into" {
		output, err := parseInsertTupleQuery(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}
		err = Controller.ProcessInsertTuple(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return "Tuple inserted successfully"

	} else if query == "list all tables"{
		output := Controller.ListAllTables()
		return output
	} else if querySplit[0] == "select" {     
		output, err := parseSelectQuery(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		tuples, err := Controller.ProcessSelect(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return fmt.Sprintf("%v", tuples)
	} else if querySplit[0] == "quit" {
		return "exiting"
	} else {
		return "Please enter a valid command"
	}
}
