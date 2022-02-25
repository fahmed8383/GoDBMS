package SQLParser

import ("fmt"
        "strings"
        "bufio"
        "os")

func ParseInput() {  
    fmt.Println("Welcome to GoDBMS")


    for true{

    //iniate buffer to read standard input delimited by newline
	reader := bufio.NewReader(os.Stdin)
	input, err1 := reader.ReadString('\n')

	if err1 != nil {
		fmt.Println("An error occured while reading input. Please try again", err1)
		return
	}

    //remove delimiter from string, convert chars to lowercase, split string by word
    input = strings.TrimSuffix(input, "\n")
    query := strings.ToLower(input) 
    querySplit := strings.Split(query, " ")

    fmt.Println(query)


    //if the user input is a 'create table' query call parseCreateTableQuery 
    if querySplit[0] == "create" && querySplit[1] == "table"{
            output, err := parseCreateTableQuery(query)
            if err != nil {
                fmt.Println(err)
                
            }else{
            fmt.Println(*output)}
        }

     //if the user input is a 'insert tuple' query call parseInsertTupleQuery    
    if querySplit[0] == "insert" && querySplit[1] == "into"{
            output, err := parseInsertTupleQuery(query)
            if err != nil {
                fmt.Println(err)
                
            }else{
            fmt.Println(*output)}
        
        }

    //if user enters 'quit' command stop taking input
    if querySplit[0] == "quit"{
        break
    }

}
    
}