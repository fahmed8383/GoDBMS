package ProcessSQLStatements

import (
	"GoDBMS/ParserStructs"
	"GoDBMS/Storage"
	"GoDBMS/StorageLock"
	"errors"
)

// ProcessCreateTable is a function that takes in a pointer to the
// CreateTableStatement struct and inserts that into the database catalog. It
// returns any errors that occur.
func ProcessCreateTable(table *ParserStructs.CreateTableStatement) (error) {

	StorageLock.AcquireCatalogLock()

	// If the table name already exists in the catalog, return an error stating so
	if Storage.TableExists(table.Name) {
		StorageLock.ReleaseCatalogLock()
		return errors.New("Table with the provided name already exist. Unable to create table.")
	}

	// Otherwise insert the table into the catalog.
	Storage.InsertTable(table)

	StorageLock.ReleaseCatalogLock()
	return nil
}
