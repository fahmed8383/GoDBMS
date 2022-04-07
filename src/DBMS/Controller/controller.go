package Controller

import (
	p "GoDBMS/Parser"
	s "GoDBMS/ProcessSQLStatements"
	"GoDBMS/Encoders"
	"os"
	"strings"
)


func InitializeCatalog() {
	// Initialize the data directory and load the catalog.
	Encoders.InitializeDirectory()
	Encoders.DecodeCatalog()
}

func StartDBMS(query string) (string) {

	// Make sure input string does not contain delimeter we will be using
	// to split strings over
	if strings.Contains(query, "?") {
		return "ERROR: Query cannot contain special character ?"
	}

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
	} else if querySplit[0] == "alter" && querySplit[1] == "table" {
		//modify table
		output, err := p.ParseModifyTable(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}
		err = s.ProcessModifyTable(output)

		if err != nil {
			return "ERROR: "+err.Error()
		}

		return "Table modified successfully"
	
	} else if query == "list all tables"{
		output := s.ListAllTables()
		return output
	} else if querySplit[0] == "select" {     
		output, err := p.ParseSelect(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		searchResult, err := s.ProcessSelect(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return searchResult
	} else if querySplit[0] == "delete" {
		output, err := p.ParseDeleteTuple(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}
		err = s.ProcessDeleteTuple(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return "Tuples deleted successfully"
	} else if querySplit[0] == "update" {
		output, err := p.ParseModifyTuple(query)
		if err != nil {
			return "ERROR: "+err.Error()
		}
		err = s.ProcessModifyTuple(output)
		if err != nil {
			return "ERROR: "+err.Error()
		}

		return "Tuples modified successfully"
	} else if querySplit[0] == "shutdown" {

		// Save the catalog before exiting.
		Encoders.EncodeCatalog()
		os.Exit(0)
		return ""

	} else {
		return "Please enter a valid command"
	}
}
