package main

import (
	"GoDBMS/SQLParser"
	"GoDBMS/Controller"
)

func main() {
	// Initialize the data directory and load the catalog.
	Controller.InitializeDirectory()
	Controller.DecodeCatalog()

	// Run the CLI to take the user input.
	SQLParser.ParseInput()
	
	// Save the catalog before exiting.
	Controller.EncodeCatalog()
}
