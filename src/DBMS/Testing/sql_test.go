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
)

func TestIncorrectCommand(t *testing.T) {
	query := "hello world" 
	err := Controller.StartDBMS(query)
	if err != "Please enter a valid command" {
		t.Errorf("ERROR: Incorrect command did not give expected error")
	}
}

func TestCreateTableMissingName(t *testing.T) {
   query := "create table (age int);"
   _, err := Parser.ParseCreateTable(query)

   if err.Error() != "Create table query has an invalid table name" {
      t.Errorf("ERROR: Create Table command missing name did not give expected error")
   }
}

func TestCreateTableMissingColumns(t *testing.T) {
	query := "create table person ();"
	_, err := Parser.ParseCreateTable(query)

   if err.Error() != "Create table statement must have a name and a type for each column" {
		t.Errorf("ERROR: Create Table command missing columns did not give expected error")
	}
}

func TestCreateTableMissingPrimaryKey(t *testing.T) {
   query := "create table person (age int);"
   _, err := Parser.ParseCreateTable(query)

   if err.Error() != "Create table statement missing a primary key" {
      t.Errorf("ERROR: Create Table command missing primary key did not give expected error")
   }
}

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

func TestProcessInsertTuple(t *testing.T) {

   Encoders.SetTestingPath()
   Storage.InitializeTables()

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

   if !Storage.TupleExists(interface{}(1), 0) || Storage.GetTuple(interface{}(1), 0).Values[1] != interface{}("Bob") {
      t.Errorf("ERROR: ProcessInsertTuple output gave the wrong tuple value, %v", err)
   }

   err = Encoders.DeleteFile("insert")
   if err != nil {
      t.Errorf("ERROR: Unable to delete file, %v", err)
   }
}

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