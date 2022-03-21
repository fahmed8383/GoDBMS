package main

import (
	"GoDBMS/Controller"
	"GoDBMS/SQLParser"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {

	// Initialize the data directory and load the catalog.
	Controller.InitializeDirectory()
	Controller.DecodeCatalog()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		body, _ := ioutil.ReadAll(r.Body)
		out := SQLParser.ParseInput(string(body))
		w.Write([]byte(out))

		if out == "DBMS Shutdown" {
			// Save the catalog before exiting.
			Controller.EncodeCatalog()
			os.Exit(0)
		}
	})

	http.ListenAndServe(":6060", nil)
}
