package Testing

import (
	"testing" 
   "strings"
	"GoDBMS/ProcessSQLStatements"
	"GoDBMS/Controller"
	"GoDBMS/Parser"
   "GoDBMS/ParserStructs"
   "GoDBMS/Storage"
   "GoDBMS/Encoders"
   "GoDBMS/StorageLock"
   "time"
)

// IP-T1
func TestIncorrectCommand(t *testing.T) {
	query := "hello world" 
	err := Controller.StartDBMS(query)
	if err != "Please enter a valid command" {
		t.Errorf("ERROR: Incorrect command did not give expected error")
	}
}

// IP-BE1-T1
func TestCreateTableMissingName(t *testing.T) {
   query := "create table (age int);"
   _, err := Parser.ParseCreateTable(query)

   if err.Error() != "Create table query has an invalid table name" {
      t.Errorf("ERROR: Create Table command missing name did not give expected error")
   }
}

// IP-BE1-T2
func TestCreateTableMissingColumns(t *testing.T) {
	query := "create table person ();"
	_, err := Parser.ParseCreateTable(query)

   if err.Error() != "Create table statement must have a name and a type for each column" {
		t.Errorf("ERROR: Create Table command missing columns did not give expected error")
	}
}

// IP-BE1-T3
func TestCreateTableMissingPrimaryKey(t *testing.T) {
   query := "create table person (age int);"
   _, err := Parser.ParseCreateTable(query)

   if err.Error() != "Create table statement missing a primary key" {
      t.Errorf("ERROR: Create Table command missing primary key did not give expected error")
   }
}

// IP-BE1-T3
func TestParseCreateTable(t *testing.T) {
   query := "create table person (id int primary key, name string, age int);"

   res, err := Parser.ParseCreateTable(query)

   if err != nil {
      t.Errorf("ERROR: Create Table command did not go through")
   }

   if res.Name != "person" {
      t.Errorf("ERROR: ProcessCreateTable output gave the wrong table name, %v", err)
   }

   if res.PrimaryKeyIndex != 0 {
      t.Errorf("ERROR: ProcessCreateTable output gave the wrong primary key index, %v", err)
   }

   if res.Columns[0].Name != "id" {
      t.Errorf("ERROR: ProcessCreateTable output gave the wrong column name, %v", err)
   }

   if res.Columns[0].Datatype != "int" {
      t.Errorf("ERROR: ProcessCreateTable output gave the wrong column datatype, %v", err)
   }

   if res.Columns[0].NotNull != true {
      t.Errorf("ERROR: ProcessCreateTable output gave the wrong column not null value, %v", err)
   }
}


// IP-BE2-T1
func TestParseDeleteTableMissingTable(t *testing.T) {
   query := "create table person (id int primary key, name string, age int);"

   _, err := Parser.ParseCreateTable(query)

   if err != nil {
      t.Errorf("ERROR: Create Table command did not go through")
   }

   query = "delete table;"

   res, err := Parser.ParseDeleteTable(query)

   if err == nil {
      t.Errorf("ERROR: ParseDeleteTable didn't give expected error, %s", res.TableName)
   }
}

// IP-BE2-T2
func TestParseDeleteTable(t *testing.T) {
   query := "create table person (id int primary key, name string, age int);"

   _, err := Parser.ParseCreateTable(query)

   if err != nil {
      t.Errorf("ERROR: Create Table command did not go through")
   }

   query = "delete table person;"

   res, err := Parser.ParseDeleteTable(query)

   if res.TableName != "person" {
      t.Errorf("ERROR: ParseDeleteTable output gave the wrong table name, %s", res.TableName)
   }
}

// IP-BE4-T1
func TestParseInsertTupleMissingTable(t *testing.T) {
   query := "insert into (id, name, age) values (1, Bob, 20);"

   _, err := Parser.ParseInsertTuple(query)

   if err == nil {
      t.Errorf("ERROR: ParseInsertTuple did not give expected error")
   }
}


// IP-BE4-T2
func TestParseInsertTupleMissingColumn(t *testing.T) {
   query := "insert into person (id) values (1, Bob);"

   _, err := Parser.ParseInsertTuple(query)

   if err == nil {
      t.Errorf("ERROR: ParseInsertTuple didn't give expected error")
   }
}

// IP-BE4-T3
func TestParseInsertTupleMissingValue(t *testing.T) {
   query := "insert into person (id, name, age) values (1, Bob);"

   _, err := Parser.ParseInsertTuple(query)

   if err == nil {
      t.Errorf("ERROR: ParseInsertTuple did not give expected error")
   }
}


// IP-BE4-T4
func TestParseInsertTuple(t *testing.T) {
   query := "insert into person (id, name, age) values (1, Bob, 20);"

   res, err := Parser.ParseInsertTuple(query)

   if err != nil {
      t.Errorf("ERROR: ParseInsertTuple did not go through")
   }

   if res.TableName != "person" {
      t.Errorf("ERROR: ParseInsertTuple output gave the wrong table name, %v", err)
   }
}

// IP-BE5-T1
// func TestParseModifyTupleMissingTable(t *testing.T) {
//    query := "update ..."

//    _, err := Parser.ParseModifyTuple(query)

//    if err == nil {
//       t.Errorf("ERROR: ParseModifyTuple did not give expected error")
//    }
// }

// IP-BE5-T2
// func TestParseModifyTuplePrimaryKey(t *testing.T) {
//    query := "update ..."

//    _, err := Parser.ParseModifyTuple(query)

//    if err == nil {
//       t.Errorf("ERROR: ParseModifyTuple did not give expected error")
//    }
// }

// IP-BE5-T3
// func TestParseModifyTupleMissingValue(t *testing.T) {
//    query := "update ..."

//    _, err := Parser.ParseModifyTuple(query)

//    if err == nil {
//       t.Errorf("ERROR: ParseModifyTuple did not give expected error")
//    }
// }

// IP-BE5-T4
// func TestParseModifyTuple(t *testing.T) {
//    query := "update ..."

//    res, err := Parser.ParseModifyTuple(query)

//    if err != nil {
//       t.Errorf("ERROR: ParseModifyTuple did not go through")
//    }

// }

// IP-BE6-T1
func TestParseDeleteTupleMissingTable(t *testing.T) {
   query := "delete from where id = 1;"

   _, err := Parser.ParseDeleteTuple(query)

   if err == nil {
      t.Errorf("ERROR: ParseDeleteTuple did not give expected error")
   }
}


// Test currently not working

// IP-BE6-T2
// func TestParseDeleteTupleMissingPrimaryKey(t *testing.T) {
//    query := "delete from person"

//    _, err := Parser.ParseDeleteTuple(query)

//    if err == nil {
//       t.Errorf("ERROR: ParseDeleteTuple did not give expected error")
//    }
// }

// IP-BE6-T3
func TestParseDeleteTuple(t *testing.T) {

   query := "delete from person where id = 1;"

   res, err := Parser.ParseDeleteTuple(query)

   if err != nil {
      t.Errorf("ERROR: ParseDeleteTuple did not go through")
   }

   if res.From != "person" {
      t.Errorf("ERROR: ParseDeleteTuple output gave the wrong table name, %v", err)
   }

   if res.Where[0] != "id" {
      t.Errorf("ERROR: ParseDeleteTuple output gave the wrong column name, %v", err)
   }

   if res.Where[2] != "1" {
      t.Errorf("ERROR: ParseDeleteTuple output gave the wrong column value, %v", err)
   }
}

// IP-BE7-T1
func TestParseSelectMissingTable(t *testing.T) {
   query := "select * from;"

   _, err := Parser.ParseSelect(query)

   if err == nil {
      t.Errorf("ERROR: ParseSelect did not give expected error")
   }
}

// IP-BE7-T2
func TestParseSelectMissingColumns(t *testing.T) {
   query := "select from person;"

   _, err := Parser.ParseSelect(query)

   if err == nil {
      t.Errorf("ERROR: ParseSelect did not give expected error")
   }
}

// IP-BE7-T3
func TestParseSelect(t *testing.T) {
   query := "select * from person;"

   res, err := Parser.ParseSelect(query)

   if err != nil {
      t.Errorf("ERROR: ParseSelect did not go through")
   }

   if res.From != "person" {
      t.Errorf("ERROR: ParseSelect output gave the wrong table name, %v", err)
   }
}

// SD-BE1-T2
func TestProcessCreateTable(t *testing.T) {

   Encoders.SetTestingPath()
   Storage.InitializeTables()

   tableStatement := ParserStructs.CreateTableStatement{"person", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   if Storage.GetTable("person") == nil {
      t.Errorf("ERROR: ProcessCreateTable did not create the table")
   }

   if Storage.GetTable("person").PrimaryKeyIndex != 0 {
      t.Errorf("ERROR: ProcessCreateTable did not set the primary key index")
   }

   if Storage.GetTable("person").Columns[0].Name != "id" {
      t.Errorf("ERROR: ProcessCreateTable did not set the column name")
   }

   if Storage.GetTable("person").Columns[0].Datatype != "int" {
      t.Errorf("ERROR: ProcessCreateTable did not set the column datatype")
   }

   if Storage.GetTable("person").Columns[0].NotNull != true {
      t.Errorf("ERROR: ProcessCreateTable did not set the column not null value")
   }
}

// SD-BE1-T3
func TestProcessCreateDuplicateTable(t *testing.T) {

   Encoders.SetTestingPath()
   Storage.InitializeTables()

   tableStatement := ParserStructs.CreateTableStatement{"person", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   tableStatement = ParserStructs.CreateTableStatement{"person", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err = ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err == nil {
      t.Errorf("ERROR: ProcessCreateTable did not give expected error")
   }
}

// SD-BE2-T2
func TestProcessDeleteTable(t *testing.T) {
   Encoders.SetTestingPath()
   Storage.InitializeTables()

   tableStatement := ParserStructs.CreateTableStatement{"person", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   deleteTableStatement := ParserStructs.DeleteTableStatement{"person"}
   err = ProcessSQLStatements.ProcessDeleteTable(&deleteTableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessDeleteTable did not go through")
   }

   if Storage.TableExists("person") {
      t.Errorf("ERROR: ProcessDeleteTable did not delete the table")
   }
}

// SD-BE2-T3
func TestProcessDeleteNonexistentTable(t *testing.T) {
   Encoders.SetTestingPath()
   Storage.InitializeTables()

   deleteTableStatement := ParserStructs.DeleteTableStatement{"person"}
   err := ProcessSQLStatements.ProcessDeleteTable(&deleteTableStatement)

   if err == nil {
      t.Errorf("ERROR: ProcessDeleteTable did not give expected error")
   }
}

// SD-BE3-T1
func TestListAllTables(t *testing.T) {

   Encoders.SetTestingPath()
   Storage.InitializeTables()

   tableStatement := ParserStructs.CreateTableStatement{"person", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   tableStatement = ParserStructs.CreateTableStatement{"person2", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err = ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   tableStatement = ParserStructs.CreateTableStatement{"person3", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err = ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   tableStatement = ParserStructs.CreateTableStatement{"person4", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err = ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   tables := ProcessSQLStatements.ListAllTables()

   if !strings.Contains(tables, "person") {
      t.Errorf("ERROR: ListAllTables does not contain all tables")
   }
   if !strings.Contains(tables, "person2") {
      t.Errorf("ERROR: ListAllTables does not contain all tables %s", "person2")
   }   
   if !strings.Contains(tables, "person3") {
      t.Errorf("ERROR: ListAllTables does not contain all tables %s", "person3")
   }   
   if !strings.Contains(tables, "person4") {
      t.Errorf("ERROR: ListAllTables does not contain all tables %s", "person4")
   }
}

// SD-BE4-T2
func TestProcessInsertTuple(t *testing.T) {

   Encoders.SetTestingPath()
   Storage.InitializeTables()
   StorageLock.InitializeLocks()

   tableStatement := ParserStructs.CreateTableStatement{"insert", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through, %v", err)
   }

   insertTuple1 := ParserStructs.InsertTupleColumn{"id", "1"}
   insertTuple2 := ParserStructs.InsertTupleColumn{"name", "Bob"}
   insertTuple3 := ParserStructs.InsertTupleColumn{"age", "20"}
   insertTupleColumns := []ParserStructs.InsertTupleColumn{insertTuple1, insertTuple2, insertTuple3}

   insertStatement := ParserStructs.InsertTupleStatement{"insert", insertTupleColumns}

   err = ProcessSQLStatements.ProcessInsertTuple(&insertStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessInsertTuple did not go through, %v", err)
   }

   heap, err := Encoders.DecodeHeap("insert")
   if err != nil {
      t.Errorf("ERROR: Decoding heap failed, %v", err)
   }

   if !heap.TupleExists(interface{}(1), 0) || heap.GetTuple(interface{}(1), 0).Values[1] != interface{}("Bob") {
      t.Errorf("ERROR: ProcessInsertTuple output gave the wrong tuple value, %v", err)
   }

   err = Encoders.DeleteFile("insert")
   if err != nil {
      t.Errorf("ERROR: Unable to delete file, %v", err)
   }
}

// SD-BE4-T3
func TestProcessInsertTupleWrongTable(t *testing.T) {

   Encoders.SetTestingPath()
   Storage.InitializeTables()
   StorageLock.InitializeLocks()
   insertTuple1 := ParserStructs.InsertTupleColumn{"id", "1"}
   insertTuple2 := ParserStructs.InsertTupleColumn{"name", "Bob"}
   insertTuple3 := ParserStructs.InsertTupleColumn{"age", "20"}
   insertTupleColumns := []ParserStructs.InsertTupleColumn{insertTuple1, insertTuple2, insertTuple3}

   insertStatement := ParserStructs.InsertTupleStatement{"insert", insertTupleColumns}

   err := ProcessSQLStatements.ProcessInsertTuple(&insertStatement)

   if err == nil {
      t.Errorf("ERROR: ProcessInsertTuple did not give expected error")
   }
}

// SD-BE5-T2
// modify the record correctly

// SD-BE5-T3
// modify the record in an incorrect table

// SD-BE5-T4
// modify the record with incorrect primary key

// SD-BE6-T2
func TestProcessDeleteTuple(t *testing.T) {

   Encoders.SetTestingPath()
   Storage.InitializeTables()
   StorageLock.InitializeLocks()

   tableStatement := ParserStructs.CreateTableStatement{"delete", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through, %v", err)
   }

   insertTuple1 := ParserStructs.InsertTupleColumn{"id", "1"}
   insertTuple2 := ParserStructs.InsertTupleColumn{"name", "Bob"}
   insertTuple3 := ParserStructs.InsertTupleColumn{"age", "20"}
   insertTupleColumns := []ParserStructs.InsertTupleColumn{insertTuple1, insertTuple2, insertTuple3}

   insertStatement := ParserStructs.InsertTupleStatement{"delete", insertTupleColumns}

   err = ProcessSQLStatements.ProcessInsertTuple(&insertStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessInsertTuple did not go through, %v", err)
   }
   
   query := "delete from delete where id = 1;"

   deleteStatement, err := Parser.ParseDeleteTuple(query)

   if err != nil {
      t.Errorf("ERROR: ParseDeleteTuple did not go through, %v", err)
   }

   err = ProcessSQLStatements.ProcessDeleteTuple(deleteStatement)
   if err != nil {
      t.Errorf("ERROR: ProcessDeleteTuple did not go through, %v", err)
   }

   heap, err := Encoders.DecodeHeap("delete")
   if err != nil {
      t.Errorf("ERROR: Decoding heap failed, %v", err)
   }

   if heap.TupleExists(interface{}(1), 0) {
      t.Errorf("ERROR: ProcessDeleteTuple failed to delete, %v", err)
   }

   err = Encoders.DeleteFile("delete")
   if err != nil {
      t.Errorf("ERROR: Unable to delete file, %v", err)
   }
}

// SD-BE6-T3
func TestProcessDeleteTupleWrongTable(t *testing.T) {
   Encoders.SetTestingPath()
   Storage.InitializeTables()
   StorageLock.InitializeLocks()

   tableStatement := ParserStructs.CreateTableStatement{"delete", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through, %v", err)
   }

   insertTuple1 := ParserStructs.InsertTupleColumn{"id", "1"}
   insertTuple2 := ParserStructs.InsertTupleColumn{"name", "Bob"}
   insertTuple3 := ParserStructs.InsertTupleColumn{"age", "20"}
   insertTupleColumns := []ParserStructs.InsertTupleColumn{insertTuple1, insertTuple2, insertTuple3}

   insertStatement := ParserStructs.InsertTupleStatement{"delete", insertTupleColumns}

   err = ProcessSQLStatements.ProcessInsertTuple(&insertStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessInsertTuple did not go through, %v", err)
   }
   
   query := "delete from wrong where id = 1;"

   deleteStatement, err := Parser.ParseDeleteTuple(query)

   if err != nil {
      t.Errorf("ERROR: ParseDeleteTuple did not go through, %v", err)
   }

   err = ProcessSQLStatements.ProcessDeleteTuple(deleteStatement)
   if err == nil {
      t.Errorf("ERROR: ProcessDeleteTuple did not give expected error")
   }

   err = Encoders.DeleteFile("delete")
   if err != nil {
      t.Errorf("ERROR: Unable to delete file, %v", err)
   }
}

// SD-BE7-T2
// select the record in an incorrect table
func TestProcessSelectWrongTable(t *testing.T) {
   Encoders.SetTestingPath()
   Storage.InitializeTables()
   StorageLock.InitializeLocks()

   tableStatement := ParserStructs.CreateTableStatement{"delete", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   err := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through, %v", err)
   }

   insertTuple1 := ParserStructs.InsertTupleColumn{"id", "1"}
   insertTuple2 := ParserStructs.InsertTupleColumn{"name", "Bob"}
   insertTuple3 := ParserStructs.InsertTupleColumn{"age", "20"}
   insertTupleColumns := []ParserStructs.InsertTupleColumn{insertTuple1, insertTuple2, insertTuple3}

   insertStatement := ParserStructs.InsertTupleStatement{"delete", insertTupleColumns}

   err = ProcessSQLStatements.ProcessInsertTuple(&insertStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessInsertTuple did not go through, %v", err)
   }
   
   query := "select wrongColumn from wrong;"

   selectStatement, err := Parser.ParseSelect(query)

   if err != nil {
      t.Errorf("ERROR: ParseSelect did not go through, %v", err)
   }

   _, err = ProcessSQLStatements.ProcessSelect(selectStatement)
   if err == nil {
      t.Errorf("ERROR: ProcessSelect did not give expected error")
   }

   err = Encoders.DeleteFile("delete")
   if err != nil {
      t.Errorf("ERROR: Unable to delete file, %v", err)
   }
} 

// SD-BE8-T2
// lock the record 

// -------------------------------- NFR ------------------------------------

func TestResponseTime(t *testing.T) {
   Encoders.SetTestingPath()
   Storage.InitializeTables()
   StorageLock.InitializeLocks()

   query := "create table person (id int primary key, name string, age int);"

   res, err := Parser.ParseCreateTable(query)
   err = ProcessSQLStatements.ProcessCreateTable(res)

   if err != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through, %v", err)
   }

   // --- Tuple 1 ---

   query = "insert into person (id, name, age) values (1, Bob, 20);"

   tuple1, err := Parser.ParseInsertTuple(query)

   if err != nil {
      t.Errorf("ERROR: ParseInsertTuple 1 did not go through")
   }

   err = ProcessSQLStatements.ProcessInsertTuple(tuple1)

   if err != nil {
      t.Errorf("ERROR: ProcessInsertTuple did not go through, %v", err)
   }

   // --- Tuple 2 ---

   query = "insert into person (id, name, age) values (2, Bob, 20);"

   tuple2, err := Parser.ParseInsertTuple(query)

   if err != nil {
      t.Errorf("ERROR: ParseInsertTuple 1 did not go through")
   }

   err = ProcessSQLStatements.ProcessInsertTuple(tuple2)

   if err != nil {
      t.Errorf("ERROR: ProcessInsertTuple did not go through, %v", err)
   }

   // --- Tuple 3 ---

   query = "insert into person (id, name, age) values (3, Bob, 20);"

   tuple3, err := Parser.ParseInsertTuple(query)

   if err != nil {
      t.Errorf("ERROR: ParseInsertTuple 1 did not go through")
   }

   err = ProcessSQLStatements.ProcessInsertTuple(tuple3)

   if err != nil {
      t.Errorf("ERROR: ProcessInsertTuple did not go through, %v", err)
   }

   // --- Select tuples --- 

   start := time.Now()

   selectTuples, err := Parser.ParseSelect("select * from person;")

   if err != nil {
      t.Errorf("ERROR: ParseSelectTuples did not go through")
   }

   _, err = ProcessSQLStatements.ProcessSelect(selectTuples)

   if err != nil {
      t.Errorf("ERROR: ProcessSelectTuples did not go through, %v", err)
   }

   elapsed := time.Since(start)

   if elapsed.Seconds() > 5 {
      t.Errorf("ERROR: ProcessSelectTuples took too long %v", elapsed.Seconds())
   }
}