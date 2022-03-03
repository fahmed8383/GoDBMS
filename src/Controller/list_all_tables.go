package Controller

import(
	"GoDBMS/Database"
	"fmt"
)

func ListAllTables() string {

	keyMap := Database.GetTablesMap()
	
	tables := ""

	for i := range *keyMap { 
		fmt.Println(i)
		tables = tables + i + " "
	}
	
	return tables

}