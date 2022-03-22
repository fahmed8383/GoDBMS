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
	// Columns is an array of insertTupleColumn representing the columns
	Columns []InsertTupleColumn
}

// insertTupleColumn holds the insert tuple query column info received from the user
type InsertTupleColumn struct {
	// Name is a string representing name of the column
	Name string
	// Value is a string representing the value of the column
	Value string
}

//SelectStatement holds info indicated by the user for a select statement
type SelectStatement struct {
	//TableName is a string representing the name of the table
	TableName string
}

//DeleteTableStatement holds info indicated by the user for a delete statement
type DeleteTableStatement struct {
	//TableName is a string representing the name of the table
	TableName string
}