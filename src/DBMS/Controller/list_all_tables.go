package Controller

import(
	"GoDBMS/Database"
)

func ListAllTables() string {

	keyMap := Database.GetTablesMap()
	
	tables := ""

	for i := range *keyMap { 
		tables = tables + i + " "
	}
	
	return tables
}