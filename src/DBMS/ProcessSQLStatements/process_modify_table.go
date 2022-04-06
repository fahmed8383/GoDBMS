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
	modifyTable := Storage.GetTable(tableParams.tableName) 

	//Check if add or drop before modifying from the catalog
	if (tableParams.DataType == "") {
		//drop statement
		//get pointer to column we want to drop
		dropColumn := modifyTable.(tableParams.columnName)
		//remove from columns array
		modifyTable.Columns[ColumnIndex[dropColumn.Name]] = nil
		//remove its index from index map
		modifyTable.ColumnsIndex[dropColumn.Name] = nil

	}else{
		//add statement
		// add the column name - index pair to the ColumnIndex map for the modified table
		modifyTable.ColumnIndex[columnName] = len(modifyTable.Columns)
		//create a struct for the new column
		newColumn := Storage.TableColumn{tableParams.columnName, tableParams.dataType, False}
		//add the new column struct to the columns array for the modified table
		modifyTable.Columns = append(modifyTable.Columns, *newColumn)
		heap, err := Encoders.DecodeHeap(table.Name)
		if err != nil {
			return nil, err
		}
		
		tuples := heap.GetHeap()
		
		newTuples := []Storage.Tuple{}
	
		// add null values to each tuple in the new column
		for _, tuple := range tuples {
			tuple.values = append(tuple.values, nil)
			primaryKey := modifyTable.PrimaryKeyIndex
			heap.ModifyTuple(interface{}(tuple[primaryKey]),primaryKey,tuple)
		}
		
	}


	/*if Encoders.FileExists(table.TableName) {
		Encoders.DeleteFile(table.TableName)
	*/}

	// Otherwise modify the table.
	//Storage.DeleteTable(table.TableName)

	StorageLock.ReleaseCatalogLock()
	
	return nil
}