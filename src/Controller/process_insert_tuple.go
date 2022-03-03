package Controller

import (
	"GoDBMS/Database"
	"errors"
	"strconv"
)

// ProcessInsertTuple is a function that handles the insertion of a tuple into
// the database. It takes an InsertTupleStatement pointer as an argument and
// returns any errors that occur
func ProcessInsertTuple(insertTuple *Database.InsertTupleStatement) (err error) {

	// Declares the table that will be used later
	var table *Database.TableSchema

	// If the table does not exist, return an error stating it
	// Otherwise, it stores the table in the variable
	if Database.TableExists(insertTuple.TableName) {
		table = Database.GetTable(insertTuple.TableName)
	} else {
		err = errors.New("Table " + insertTuple.TableName + " does not exist")
		return err
	}

	// Deserializes the heap from the file and loads it
	// An error is returned if one occurs
	err = DecodeHeap(table.Name)
	if err != nil {
		return err
	}

	// The primary key of the tuple as a string
	tupleKeyString := insertTuple.Columns[table.PrimaryKeyIndex].Value

	// Declaring tupleKey as an interface for comparisons
	var tupleKey interface{}

	// If the primary key is supposed to be an int, then it will be converted to an int
	// and assigns it to the interface
	if table.Columns[table.PrimaryKeyIndex].Datatype == "int" {
		tupleKey, err = strconv.Atoi(tupleKeyString)
		if err != nil {
			return err
		}
		// Otherwise it assigns the string value to the interface
	} else {
		tupleKey = tupleKeyString
	}

	// Utilizes the TupleExists function to check if the tuple already exists
	// If so, it returns an error stating so
	if Database.TupleExists(tupleKey, table.PrimaryKeyIndex) {
		err = errors.New("Tuple with key " + tupleKeyString + " already exists")
		return err
	}

	// Creates a new array of type InsertTupleColumn to store the values of the tuple
	// It uses make to create an array with dynamic size using the length of columns
	// in the table schema
	finalColumns := make([]Database.InsertTupleColumn, len(table.Columns))

	// Iterates through the columns of the InsertTupleStatement query and adds
	// them to the finalColumns array in their corresponding index
	// The index is determined by the hash map ColumnIndex in the table schema
	for _, col := range insertTuple.Columns {
		// If the column does not exist in the table schema, exists would be false
		// Which would then be used to return an error
		index, exists := table.ColumnIndex[col.Name]
		if exists != true {
			err = errors.New("Column " + col.Name + " does not exist")
			return err
		}
		// Stores a new InsertTupleColumn in the finalColumns array at the index
		finalColumns[index] = Database.InsertTupleColumn{col.Name, col.Value}
	}

	// Iterates through finalColumns and checks for missing columns/null indices
	// If the column is empty and it can't be Null, it returns an error
	for i, col := range finalColumns {
		if col.Name == "" {
			if table.Columns[i].NotNull {
				err = errors.New("Column " + table.Columns[i].Name + " canot be null")
				return err
				// If the column can be null, it assigns the name attribute
			} else {
				finalColumns[i].Name = table.Columns[i].Name
			}

		}
	}

	// Replaces the values of the insertTuple with the finalColumns array
	insertTuple.Columns = finalColumns

	// Creates a new Tuple pointer using the CreateTuple function
	tuple, err := Database.CreateTuple(table, insertTuple)

	// Inserts the tuple into the heap
	Database.InsertTuple(tuple)

	// Encodes the heap and saves it to the file
	err = EncodeHeap(table.Name)
	if err != nil {
		return err
	}

	return nil
}
