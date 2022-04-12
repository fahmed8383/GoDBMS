package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net/http"
	"bytes"
	"io/ioutil"
)

// Main function to run the CLI
func main() {
	fmt.Println("Welcome to GoDBMS")
	fmt.Println("Please enter a query")

	for true {

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

		if query == "quit" {
			break
		}
		
		resp, err := http.Post("http://127.0.0.1:6060/", "application/json", bytes.NewBuffer([]byte(query)))

		// If the server is shutdown sometimes the response is corrupted which
		// causes an error. Thus break instead of checking for error on shutdown
		if query == "shutdown" {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(string(body))
	}
}