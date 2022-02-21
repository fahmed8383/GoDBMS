package main 

import (  
	"fmt"
    "GoDBMS/SQLParser"
	"GoDBMS/Database"
	"GoDBMS/Modules"
)

func main() {  
    fmt.Println("Hello World")
	SQLParser.Print()
	Database.Print()
	Modules.Print()
}