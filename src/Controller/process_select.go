package Controller

import (
	"GoDBMS/Database"
	"errors"
)

func ProcessSelect(selectStatement *Database.SelectStatement) ([]Database.Tuple, error) {
	var table *Database.TableSchema
	// check if table exists, if not, return error
	if Database.TableExists(selectStatement.TableName) {
		table = Database.GetTable(selectStatement.TableName)
	} else {
		return nil, errors.New("Table from select statement does not exist")
	}

	// load heap, decode heap
	err := DecodeHeap(table.Name)
	if err != nil {
		return nil, err
	}
	heap := Database.GetHeap()
	
	tuples := []Database.Tuple{}

	// Prints tuples in heap
	for _, tuple := range *heap {
		tuples = append(tuples, *tuple)
	}

	return tuples, nil
}