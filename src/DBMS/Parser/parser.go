package Parser

import (
	"GoDBMS/ParserStructs"
	child "GoDBMS/Parser/Statements"
)

// ParseCreateTable is a interface to pass on and return the values from
// Statements.ParseCreateTable.
func ParseCreateTable(query string) (*ParserStructs.CreateTableStatement, error) {
	return child.ParseCreateTable(query)
}

// ParseInsertTuple is a interface to pass on and return the values from
// Statements.ParseInsertTuple.
func ParseInsertTuple(query string) (*ParserStructs.InsertTupleStatement, error) {
	return child.ParseInsertTuple(query)
}

// ParseSelect is a interface to pass on and return the values from
// Statements.ParseSelect.
func ParseSelect(query string) (*ParserStructs.SelectStatement, error) {
	return child.ParseSelect(query)
}

// ParseDeleteTable is a interface to pass on and return the values from
// Statements.ParseDeleteTable.
func ParseDeleteTable(query string) (*ParserStructs.DeleteTableStatement, error) {
	return child.ParseDeleteTable(query)
}

// ParseDeleteTuple is a interface to pass on and return the values from
// Statements.ParseDeleteTuple.
func ParseDeleteTuple(query string) (*ParserStructs.DeleteTupleStatement, error) {
	return child.ParseDeleteTuple(query)
}

// ParseModifyTable is a interface to pass on and return the values from
// Statements.ParseModifyTable.
func ParseModifyTable(query string) (*ParserStructs.ModifyTableStatement, error) {
	return child.ParseModifyTable(query)
}

// ParseModifyTuple is a interface to pass on and return the values from
// Statements.ParseModifyTuple.
func ParseModifyTuple(query string) (*ParserStructs.ModifyTupleStatement, error) {
	return child.ParseModifyTuple(query)
}
