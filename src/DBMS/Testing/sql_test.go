package Testing

import (
	"testing" 
	// "fmt"
   "strings"
	"GoDBMS/ProcessSQLStatements"
	"GoDBMS/Controller"
	"GoDBMS/Parser"
   "GoDBMS/ParserStructs"
   "GoDBMS/Storage"
   "GoDBMS/Encoders"
   // "errors"
)

func TestIncorrectCommand(t *testing.T) {
	query := "hello world" 
	err := Controller.StartDBMS(query)
	// fmt.Println(err)

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

func TestProcessCreateTable(t *testing.T) {

   Encoders.SetTestingPath()
   Storage.InitializeTables()

   tableStatement := ParserStructs.CreateTableStatement{"person", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   errCT := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if errCT != nil {
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

func TestProcessInsertTuple(t *testing.T) {
   // query := "insert into person (id, name, age) values (1, Bob, 20);"

   Encoders.SetTestingPath()
   Storage.InitializeTables()

   tableStatement := ParserStructs.CreateTableStatement{"insert", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   errCT := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if errCT != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through, %v", errCT)
   }

   insertTuple1 := ParserStructs.InsertTupleColumn{"id", "5"}
   insertTuple2 := ParserStructs.InsertTupleColumn{"name", "Bob"}
   insertTuple3 := ParserStructs.InsertTupleColumn{"age", "20"}
   insertTupleColumns := []ParserStructs.InsertTupleColumn{insertTuple1, insertTuple2, insertTuple3}

   insertStatement := ParserStructs.InsertTupleStatement{"insert", insertTupleColumns}

   err := ProcessSQLStatements.ProcessInsertTuple(&insertStatement)

   if err != nil {
      t.Errorf("ERROR: ProcessInsertTuple did not go through, %v", err)
   }

   var tupleKey interface{}
   tupleKey = 1


   if Storage.GetTuple(tupleKey, 0).Values[0] != interface{}(5) {
      t.Errorf("ERROR: ProcessInsertTuple output gave the wrong tuple value, %v", err)
   }

}

func TestListAllTables(t *testing.T) {

   Encoders.SetTestingPath()
   Storage.InitializeTables()

   tableStatement := ParserStructs.CreateTableStatement{"person", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   errCT := ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if errCT != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   tableStatement = ParserStructs.CreateTableStatement{"person2", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   errCT = ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if errCT != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   tableStatement = ParserStructs.CreateTableStatement{"person3", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   errCT = ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if errCT != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   tableStatement = ParserStructs.CreateTableStatement{"person4", 0, []ParserStructs.CreateTableColumn{ParserStructs.CreateTableColumn{"id", "int", true}, ParserStructs.CreateTableColumn{"name", "string", false}, ParserStructs.CreateTableColumn{"age", "int", false}}}
   errCT = ProcessSQLStatements.ProcessCreateTable(&tableStatement)

   if errCT != nil {
      t.Errorf("ERROR: ProcessCreateTable did not go through")
   }

   tables := ProcessSQLStatements.ListAllTables()

   // if len(tables) != 4 {
   //    t.Errorf("ERROR: ListAllTables did not return the correct number of tables, %v", err)
   // }

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
   // tablesMap := make(map[string]bool)


   // if tables != "person person2 person3 person4" {
   //    t.Errorf("ERROR: ListAllTables did not list all the tables: %s", tables)
   // }

}