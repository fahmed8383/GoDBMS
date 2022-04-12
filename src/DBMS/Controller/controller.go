package Controller

import (
	p "GoDBMS/Parser"
	s "GoDBMS/ProcessSQLStatements"
	"GoDBMS/Encoders"
	"os"
	"strings"
)

// InitializeCatalog is a function to intialize the data file directory
// and decode the catalog file to load it into memory.
func InitializeCatalog() {
	// Initialize the data directory and load the catalog.
	Encoders.InitializeDirectory()
	Encoders.DecodeCatalog()
}

// StartDBMS is a function that takes in a string query, passes it to the
// appropriate parser and controller, and then returns the output of the query
// as a string.
func StartDBMS(query string) (string) {

	// Make sure input string does not contain delimeter we will be using
	// to split strings over
	if strings.Contains(query, "?") {
		return "ERROR: Query cannot contain special character ?"
	}

	querySplit := strings.Split(query, " ")

	// Look through input query's words to see what type of SQL command it is.
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
