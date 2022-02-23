package main

import (
	"GoDBMS/Database"
	"GoDBMS/Modules"
	"GoDBMS/SQLParser"
	"fmt"
)

func main() {
	fmt.Println("I am the main file")
	SQLParser.Print()
	Database.Print()
	Modules.Print()
}
