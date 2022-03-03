package Controller

import (
	"GoDBMS/Database"
	"fmt"
)

func ProcessInsertTuple(insertTuple *Database.InsertTupleStatement) (err error) {
	var table *Database.TableSchema

	// check if table exists, if not, return error
	if Database.TableExists(insertTuple.TableName) {
		table = Database.GetTable(insertTuple.TableName)
	} else {
		err = fmt.Errorf("Table %s does not exist", insertTuple.TableName)
		return err
	}

	fmt.Println("Table found", table.Columns)

	// load heap, decode heap
	err = DecodeHeap(table.Name)
	if err != nil {
		return err
	}
	heap := Database.GetHeap()

	// Prints tuples in heap
	for _, tuple := range *heap {
		fmt.Println("Tuple", tuple)
	}

	tupleKey := insertTuple.Columns[table.PrimaryKeyIndex].Value
	fmt.Println(Database.TupleExists(tupleKey, table.PrimaryKeyIndex), tupleKey, table.PrimaryKeyIndex)
	// check if tuple with primary key already exists
	if Database.TupleExists(tupleKey, table.PrimaryKeyIndex) {
		err = fmt.Errorf("Tuple with key %s already exists", tupleKey)
		return err
	}

	// not null validity
	// columns match table and values

	// create tuple
	// insert tuple into table using heap

	tuple, err := Database.CreateTuple(table, insertTuple)

	Database.InsertTuple(tuple)

	// save heap, encode heap
	err = EncodeHeap(table.Name)
	if err != nil {
		return err
	}

	return nil
}
