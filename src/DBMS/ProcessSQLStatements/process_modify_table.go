package ProcessSQLStatements

import (
	"GoDBMS/ParserStructs"
	"GoDBMS/Storage"
	"GoDBMS/Encoders"
	"GoDBMS/StorageLock"
	"errors"
)

func ProcessModifyTable(tableParams *ParserStructs.ModifyTableStatement) (error) {

	// If the table name does not exists in the catalog, return an error stating so
	if !(Storage.TableExists(tableParams.TableName)) {
		return errors.New("Can not modify table as it does not exist.")
	}

	StorageLock.AcquireCatalogLock()

	//get pointer to the table we want to modify from catalog
	modifyTable := Storage.GetTable(tableParams.TableName) 

	//Check if add or drop before modifying from the catalog
	if (tableParams.DataType == "") {
		// drop statement

		// get index to column we want to drop and ensure that the column exists
		dropColumnIndex, exists := modifyTable.ColumnIndex[tableParams.ColumnName]
		if !exists {
			StorageLock.ReleaseCatalogLock()
			return errors.New("Cannot delete columns " + tableParams.ColumnName + " as it does not exist")
		}

		// Check to make sure that we are not deleting the primary key
		if dropColumnIndex == modifyTable.PrimaryKeyIndex {
			StorageLock.ReleaseCatalogLock()
			return errors.New("Cannot delete columns " + tableParams.ColumnName + " as it is the primary key column")
		}

		// remove from columns array
		modifyTable.Columns = append(modifyTable.Columns[:dropColumnIndex], modifyTable.Columns[dropColumnIndex+1:]...)
		// remove its index from index map
		// modifyTable.ColumnIndex[tableParams.ColumnName] = nil
		delete(modifyTable.ColumnIndex, tableParams.ColumnName)

		// shift indices left after dropped column by 1 
		for index := range modifyTable.ColumnIndex {
			if modifyTable.ColumnIndex[index] > dropColumnIndex {
				modifyTable.ColumnIndex[index]--
			}
		}

		// Done modifying catalog
		StorageLock.ReleaseCatalogLock()

		StorageLock.AcquireTableLock(modifyTable.Name)
		heap, err := Encoders.DecodeHeap(modifyTable.Name)
		if err != nil {
			StorageLock.ReleaseTableLock(modifyTable.Name)
			return err
		}
		
		tuples := heap.GetHeap()
	
		// remove column value from each tuple in the new column
		for _, tuple := range tuples {
			tuple.Values = append(tuple.Values[:dropColumnIndex], tuple.Values[dropColumnIndex+1:]...)
		}

		// Encodes the heap and saves it to the file
		err = Encoders.EncodeHeap(modifyTable.Name, heap)
		if err != nil {
			StorageLock.ReleaseTableLock(modifyTable.Name)
			return err
		}

	} else {
		// add statement

		// Check to make sure we are not adding a column name that already exists
		_, exists := modifyTable.ColumnIndex[tableParams.ColumnName]
		if exists {
			StorageLock.ReleaseCatalogLock()
			return errors.New("Cannot add column " + tableParams.ColumnName + " as it already exists")
		}
		
		// add the column name - index pair to the ColumnIndex map for the modified table
		modifyTable.ColumnIndex[tableParams.ColumnName] = len(modifyTable.Columns)
		//create a struct for the new column
		newColumn := Storage.TableColumn{tableParams.ColumnName, tableParams.DataType, false}
		//add the new column struct to the columns array for the modified table
		modifyTable.Columns = append(modifyTable.Columns, newColumn)

		// Done modifying catalog
		StorageLock.ReleaseCatalogLock()

		StorageLock.AcquireTableLock(modifyTable.Name)
		heap, err := Encoders.DecodeHeap(modifyTable.Name)
		if err != nil {
			StorageLock.ReleaseTableLock(modifyTable.Name)
			return err
		}
		
		tuples := heap.GetHeap()
	
		// add null values to each tuple in the new column
		for _, tuple := range tuples {
			tuple.Values = append(tuple.Values, nil)
		}
			
		// Encodes the heap and saves it to the file
		err = Encoders.EncodeHeap(modifyTable.Name, heap)
		if err != nil {
			StorageLock.ReleaseTableLock(modifyTable.Name)
			return err
		}
	}

	StorageLock.ReleaseTableLock(modifyTable.Name)
	return nil
}