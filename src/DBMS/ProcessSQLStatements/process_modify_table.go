package ProcessSQLStatements

import (
	"GoDBMS/ParserStructs"
	"GoDBMS/Storage"
	"GoDBMS/Encoders"
	"GoDBMS/StorageLock"
	"errors"
)

func ProcessModifyTable(tableParams *ParserStructs.ModifyTableStatement) (error) {

	StorageLock.AcquireCatalogLock()

	// If the table name does not exists in the catalog, return an error stating so
	if !(Storage.TableExists(tableParams.TableName)) {
		return errors.New("Can not modify table as it does not exist.")
	}
	//get pointer to the table we want to modify from catalog
	modifyTable := Storage.GetTable(tableParams.TableName) 

	//Check if add or drop before modifying from the catalog
	if (tableParams.DataType == "") {
		// drop statement

		// get index to column we want to drop
		dropColumnIndex := modifyTable.ColumnIndex[tableParams.ColumnName]

		// remove from columns array
		modifyTable.Columns = append(modifyTable.Columns[:dropColumnIndex], modifyTable.Columns[dropColumnIndex+1:]...)
		// remove its index from index map
		// modifyTable.ColumnIndex[tableParams.ColumnName] = nil
		delete(modifyTable.ColumnIndex, tableParams.ColumnName)

		// shift indices after column by 1 
		for index := range modifyTable.ColumnIndex {
			if modifyTable.ColumnIndex[index] > dropColumnIndex {
				modifyTable.ColumnIndex[index]--
			}
		}

		heap, err := Encoders.DecodeHeap(modifyTable.Name)
		if err != nil {
			return err
		}
		
		tuples := heap.GetHeap()
	
		// remove column value to each tuple in the new column
		for _, tuple := range tuples {
			// tuple.Values = append(tuple.Values, nil)
			tuple.Values = append(tuple.Values[:dropColumnIndex], tuple.Values[dropColumnIndex+1:]...)
			primaryKey := modifyTable.PrimaryKeyIndex
			tupleKey := tuple.Values[primaryKey]

			heap.ModifyTuple(interface{}(tupleKey), primaryKey, tuple)
		}

		// Encodes the heap and saves it to the file
		err = Encoders.EncodeHeap(modifyTable.Name, heap)
		if err != nil {
			return err
		}
	} else {
		// add statement
		
		// add the column name - index pair to the ColumnIndex map for the modified table
		modifyTable.ColumnIndex[tableParams.ColumnName] = len(modifyTable.Columns)
		//create a struct for the new column
		newColumn := Storage.TableColumn{tableParams.ColumnName, tableParams.DataType, false}
		//add the new column struct to the columns array for the modified table
		modifyTable.Columns = append(modifyTable.Columns, newColumn)
		heap, err := Encoders.DecodeHeap(modifyTable.Name)
		if err != nil {
			return err
		}
		
		tuples := heap.GetHeap()
	
		// add null values to each tuple in the new column
		for _, tuple := range tuples {
			tuple.Values = append(tuple.Values, nil)
			primaryKey := modifyTable.PrimaryKeyIndex
			tupleKey := tuple.Values[primaryKey]

			heap.ModifyTuple(interface{}(tupleKey), primaryKey, tuple)
		}
			
		// Encodes the heap and saves it to the file
		err = Encoders.EncodeHeap(modifyTable.Name, heap)
		if err != nil {
			return err
		}
	}

	StorageLock.ReleaseCatalogLock()
	
	return nil
}