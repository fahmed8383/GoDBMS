package main

import (
	p "GoDBMS/Parser"
	s "GoDBMS/ProcessSQLStatements"
	"GoDBMS/Encoders"
	"net/http"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
)

func main() {

	InitializeCatalog()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		body, _ := ioutil.ReadAll(r.Body)
		out := StartDBMS(string(body))
		w.Write([]byte(out))
	})

	http.ListenAndServe(":6060", nil)
}

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
