package ProcessSQLStatements

import(
	"GoDBMS/Storage"
)

func ListAllTables() string {

	keyMap := Storage.GetTablesMap()
	
	tables := ""

	for i := range *keyMap { 
		tables = tables + i + " "
	}
	
	return tables
}