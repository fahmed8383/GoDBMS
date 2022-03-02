package SQLParser

import (
    "GoDBMS/Controller"
    "fmt"
    "strings"
    "bufio"
    "os"
)

func ParseInput() {

    fmt.Println("Welcome to GoDBMS")
    fmt.Println("Please enter a query")

    for true{

        fmt.Println("")

        //iniate buffer to read standard input delimited by newline
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')

        if err != nil {
            fmt.Println("An error occured while reading input. Please try again", err)
            return
        }

        //remove delimiter from string, convert chars to lowercase, split string by word
        //input = strings.TrimSuffix(input, "\n")
        input = strings.Replace(input, "\r", "", -1)
        input = strings.Replace(input, "\n", "", -1)
        query := strings.ToLower(input) 
        querySplit := strings.Split(query, " ")

        //if the user input is a 'create table' query call parseCreateTableQuery 
        if querySplit[0] == "create" && querySplit[1] == "table"{
            output, err := parseCreateTableQuery(query)
            if err != nil {
                fmt.Print("ERROR: ")
                fmt.Println(err)
                continue
            }

            err = Controller.ProcessCreateTable(output)
            if err != nil {
                fmt.Print("ERROR: ")
                fmt.Println(err)
                continue
            }

            fmt.Println("Table created successfully")

        } else if querySplit[0] == "insert" && querySplit[1] == "into"{
            output, err := parseInsertTupleQuery(query)
            if err != nil {
                fmt.Print("ERROR: ")
                fmt.Println(err)
            } else {
                fmt.Println(*output)
            }
        } else if querySplit[0] == "quit"{
            break
        } else {
            fmt.Println("Please enter a valid command")
        }
    }
}