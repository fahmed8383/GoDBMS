package ProcessSQLStatements

import (
	"GoDBMS/ParserStructs"
	"GoDBMS/Storage"
	"GoDBMS/Encoders"
	"errors"
)

func ProcessSelect(selectStatement *ParserStructs.SelectStatement) ([]Storage.Tuple, error) {
	var table *Storage.TableSchema
	// check if table exists, if not, return error
	if Storage.TableExists(selectStatement.TableName) {
		table = Storage.GetTable(selectStatement.TableName)
	} else {
		return nil, errors.New("Table from select statement does not exist")
	}

	// load heap, decode heap
	err := Encoders.DecodeHeap(table.Name)
	if err != nil {
		return nil, err
	}
	heap := Storage.GetHeap()
	
	tuples := []Storage.Tuple{}

	// Prints tuples in heap
	for _, tuple := range *heap {
		tuples = append(tuples, *tuple)
	}

	return tuples, nil
}