package ProcessSQLStatements

import (
	"GoDBMS/ParserStructs"
	"GoDBMS/Storage"
	"errors"
)

// ProcessDeleteTable is a function that takes in a pointer to the
// DeleteTableStatement struct and deletes that from the database catalog. It
// returns any errors that occur.
func ProcessDeleteTable(table *ParserStructs.DeleteTableStatement) (error) {

	// If the table name does not exists in the catalog, return an error stating so
	if !(Storage.TableExists(table.TableName)) {
		return errors.New("Can not delete table as it does not exist.")
	}

	// Otherwise delete the table from the catalog.
	Storage.DeleteTable(table.TableName)
	return nil
}