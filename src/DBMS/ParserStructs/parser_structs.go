package ParserStructs

// CreateTableStatement holds the create table query info received from the
// user.
type CreateTableStatement struct {
	// Name is a string representing name of the table.
	Name string
	// PrimaryKeyIndex is an int representing the index of the primary key 
	// column.
	PrimaryKeyIndex int
	// Columns is an array of CreateTableColumn structs representing the 
	// columns.
	Columns []CreateTableColumn
}

// CreateTableColumn holds the create table query column info received from the
// user.
type CreateTableColumn struct {
	// Name is a string representing name of the column.
	Name string
	// Name is a string representing datatype of the column.
	Datatype string
	// NotNull is a boolean representing whether the entries into the column
	// can be null or not.
	NotNull bool
}

// InsertTupleStatement holds the insert tuple query info received from the user
type InsertTupleStatement struct {
	// TableName is a string representing name of the table
	TableName string
	// Columns is an array of InsertTupleColumn representing the columns
	Columns []InsertTupleColumn
}

// InsertTupleColumn holds the insert tuple query column info received from the user
type InsertTupleColumn struct {
	// Name is a string representing name of the column
	Name string
	// Value is a string representing the value of the column
	Value string
}

// SelectStatement holds info indicated by the user for a select statement
type SelectStatement struct {
	// Select is an array that holds all the select parameters
	Select []string
	// From is an array that holds the table name
	From string
	// Where is an array that holds the where column, comparator, and value
	Where []string
}

// DeleteTableStatement holds info indicated by the user for a delete statement
type DeleteTableStatement struct {
	//TableName is a string representing the name of the table
	TableName string
}

// DeleteTupleStatement holds the delete tuple query info received from the user
type DeleteTupleStatement struct {
	// From is an array that holds the table name
	From string
	// Where is an array that holds the where column, comparator, and value
	Where []string
}

// ModifyTableStatement holds info indicated by the user for a modify table statement
type ModifyTableStatement struct {
	// TableName is a string representing the name of the table to be modified
	TableName string
	// ColumnName is a string representing the name of the column to be modified
	ColumnName string
	// Datatype is a string representing the datatype of the new column to be added
	DataType string
}

// ModifyTupleStatement holds the modify tuple query info received from the user
type ModifyTupleStatement struct {
	// TableName is a string representing the name of the table to be modified
	TableName string
	// Set is a 2D array of strings where the first index in in the inner array
	// represents the column to be modifed, and the second index represents
	// the new value
	Set [][2]string
	// Where is an array that holds the where column, comparator, and value
	Where []string
}