package Database

import (
	"strings"
	"GoDBMS/SQLParser"
)

// Create a global variable tables hashmap for this package. This variable
// stores the mapping a table name and its specific table struct pointer.
var tables map[string]*TableSchema

// TableSchema is a struct that holds the information about the table's schema
type TableSchema struct {
	// Name is a string representing name of the table.
	Name string
	// PrimaryKeyIndex is an int representing the index of the primary key 
	// column.
	PrimaryKeyIndex int
	// ColumnIndex is a hashmap mapping the name of a column to its index in
	// the table.
	ColumnIndex map[string]int
	// Columns is an array of tableColumn representing the columns of the table.
	Columns []tableColumn
}

// tableColumn is a struct that holds the information about the table's columns.
type tableColumn struct {
	// Name is a string representing name of the column.
	Name string
	// Name is a string representing datatype of the column.
	Datatype string
	// NotNull is a boolean representing whether the entries into the column
	// can be null or not.
	NotNull bool
}

// InitializeTables is a function to initialize the tables map global variable.
func InitializeTables(){
	tables = make(map[string]*TableSchema)
}

// LoadTablesMap is a function to load the table name to TableSchema map
// to the tables map global variable.
func LoadTablesMap(tablesMap *map[string]*TableSchema){
	tables = *tablesMap
}

// TableExists is a function that takes a table name string input to check if
// the table name key exists in the tables map. If the table name exists, the
// function returns true, else it returns false.
func TableExists(tableName string) (bool){
	tableName = strings.ToLower(tableName)
	_, exists := tables[tableName]
	return exists
}

// InsertTable is a function to create a TableSchema struct pointer from a
// SQLParser.CreateTableStatement pointer and saves the TableSchema struct
// pointer to the tables map.
func InsertTable(tableInfo *SQLParser.CreateTableStatement){
	// Create a table column mapping and array.
	columnMap := make(map[string]int)
	columnsArray := []tableColumn{}

	// For each column create a mapping from the column name to its index and
	// create and add a tableColumn struct to the columnsArray.
	for i, column := range tableInfo.Columns {
		columnMap[column.Name] = i
		newColumn := tableColumn{column.Name, column.Datatype, column.NotNull}
		columnsArray = append(columnsArray, newColumn)
	}

	// Create the new TableSchema struct and save it to the tables map.
	newTable := TableSchema{tableInfo.Name, tableInfo.PrimaryKeyIndex, columnMap, columnsArray}
	tables[tableInfo.Name] = &newTable
}

// GetTable is a function that takes a table name string input to get and return
// the TableSchema struct pointer that the table name is mapped to in tables.
func GetTable(tableName string) (*TableSchema){
	tableName = strings.ToLower(tableName)
	return tables[tableName]
}

// DeleteTable is a function that takes a table name string input and deletes
// the TableSchema struct pointer mapped to that name.
func DeleteTable(tableName string) {
	tableName = strings.ToLower(tableName)
	delete(tables, tableName)
}

// GetTablesMap is a function that returns a pointer to the tables global variable 
func GetTablesMap() (*map[string]*TableSchema){
	return &tables
}