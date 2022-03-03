package Controller

import (
	"GoDBMS/Database"
	"errors"
	"fmt"
	"strconv"
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

	tupleKeyString := insertTuple.Columns[table.PrimaryKeyIndex].Value

    var tupleKey interface{}

    if table.Columns[table.PrimaryKeyIndex].Datatype == "int" {
        tupleKey, err = strconv.Atoi(tupleKeyString)
        if err != nil {
            return err
        }
    } else {
        tupleKey = tupleKeyString
    }

	// check if tuple with primary key already exists
	if Database.TupleExists(tupleKey, table.PrimaryKeyIndex) {
		err = fmt.Errorf("Tuple with key %s already exists", tupleKey)
		return err
	}

	finalColumns := make([]Database.InsertTupleColumn, len(table.Columns))

	for _, col := range insertTuple.Columns {
		index, exists := table.ColumnIndex[col.Name]
		if exists != true {
			err = errors.New("Column " + col.Name + " does not exist")
			return err
		}
		finalColumns[index] = Database.InsertTupleColumn{col.Name, col.Value}
	}

	for i, col := range finalColumns {
		if col.Name == "" {
			if table.Columns[i].NotNull {
				// fmt.Println(table.Columns[i])
				err = errors.New("Column " + table.Columns[i].Name + " canot be null")
				return err
			} else {
				finalColumns[i].Name = table.Columns[i].Name
			}

		}
	}

	insertTuple.Columns = finalColumns

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
