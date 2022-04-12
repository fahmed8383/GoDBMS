package ProcessSQLStatements

import(
	"GoDBMS/Storage"
)

// ListAllTables is a function that loops through all entries in the catalog
// and returns a string with the name of all the tables in the current catalog.
func ListAllTables() string {

	keyMap := Storage.GetTablesMap()
	
	tables := ""

	for i := range *keyMap { 
		tables = tables + i + " "
	}
	
	return tables
}