package Controller

import (
	p "GoDBMS/Parser"
	s "GoDBMS/ProcessSQLStatements"
	"GoDBMS/Encoders"
	"os"
	"strings"
	"fmt"
)


func InitializeCatalog() {
	// Initialize the data directory and load the catalog.
	Encoders.InitializeDirectory()
	Encoders.DecodeCatalog()
}

func StartDBMS(query string) (string) {

	querySplit := strings.Split(query, " ")

	// If the user input is a 'create table' query call ParseCreateTable
	if querySplit[0] == "create" && querySplit[1] == "table" {
		output, err := p.ParseCreateTable(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		err = s.ProcessCreateTable(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return "Table created successfully"

	} else if querySplit[0] == "insert" && querySplit[1] == "into" {
		output, err := p.ParseInsertTuple(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}
		err = s.ProcessInsertTuple(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return "Tuple inserted successfully"
	} else if querySplit[0] == "drop" && querySplit[1] == "table" {
		output, err := p.ParseDeleteTable(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}
		err = s.ProcessDeleteTable(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return "Table deleted successfully"
	} else if querySplit[1] == "alter" && querySplit[1] == "table" {
		//modify table
	
	} else if query == "list all tables"{
		output := s.ListAllTables()
		return output
	} else if querySplit[0] == "select" {     
		output, err := p.ParseSelect(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		tuples, err := s.ProcessSelect(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return fmt.Sprintf("%v", tuples)
	} else if querySplit[0] == "shutdown" {

		// Save the catalog before exiting.
		Encoders.EncodeCatalog()
		os.Exit(0)
		return ""

	} else {
		return "Please enter a valid command"
	}
}
