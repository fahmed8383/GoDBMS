package Controller

import (
	"GoDBMS/Database"
	"errors"
)

// ProcessCreateTable is a function that takes in a pointer to the
// CreateTableStatement struct and inserts that into the database catalog. It
// returns any errors that occur.
func ProcessCreateTable(table *Database.CreateTableStatement) (error) {

	// If the table name already exists in the catalog, return an error stating so
	if Database.TableExists(table.Name) {
		return errors.New("Table with the provided name already exist. Unable to create table.")
	}

	// Otherwise insert the table into the catalog.
	Database.InsertTable(table)
	return nil
}
