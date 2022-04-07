package Parser

import (
	"GoDBMS/ParserStructs"
	child "GoDBMS/Parser/Statements"
)

func ParseCreateTable(query string) (*ParserStructs.CreateTableStatement, error) {
	return child.ParseCreateTable(query)
}

func ParseInsertTuple(query string) (*ParserStructs.InsertTupleStatement, error) {
	return child.ParseInsertTuple(query)
}

func ParseSelect(query string) (*ParserStructs.SelectStatement, error) {
	return child.ParseSelect(query)
}

func ParseDeleteTable(query string) (*ParserStructs.DeleteTableStatement, error) {
	return child.ParseDeleteTable(query)
}

func ParseDeleteTuple(query string) (*ParserStructs.DeleteTupleStatement, error) {
	return child.ParseDeleteTuple(query)
}

func ParseModifyTable(query string) (*ParserStructs.ModifyTableStatement, error) {
	return child.ParseModifyTable(query)
}

func ParseModifyTuple(query string) (*ParserStructs.ModifyTupleStatement, error) {
	return child.ParseModifyTuple(query)
}
