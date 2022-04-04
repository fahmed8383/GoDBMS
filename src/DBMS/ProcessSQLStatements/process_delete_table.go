package ProcessSQLStatements

import (
	"GoDBMS/ParserStructs"
	"GoDBMS/Storage"
	"GoDBMS/Encoders"
	"GoDBMS/StorageLock"
	"errors"
)

// ProcessDeleteTable is a function that takes in a pointer to the
// DeleteTableStatement struct and deletes that from the database catalog. It
// returns any errors that occur.
func ProcessDeleteTable(table *ParserStructs.DeleteTableStatement) (error) {

	StorageLock.AcquireCatalogLock()

	// If the table name does not exists in the catalog, return an error stating so
	if !(Storage.TableExists(table.TableName)) {
		return errors.New("Can not delete table as it does not exist.")
	}

	if Encoders.FileExists(table.TableName) {
		Encoders.DeleteFile(table.TableName)
	}

	// Otherwise delete the table from the catalog.
	Storage.DeleteTable(table.TableName)

	StorageLock.ReleaseCatalogLock()
	
	return nil
}