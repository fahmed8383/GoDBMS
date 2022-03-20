package main

import (
	"GoDBMS/Controller"
	"GoDBMS/SQLParser"
	"net/http"
	"io/ioutil"
)

func main() {

	// Initialize the data directory and load the catalog.
	Controller.InitializeDirectory()
	Controller.DecodeCatalog()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		body, _ := ioutil.ReadAll(r.Body)
		out := SQLParser.ParseInput(string(body))
		w.Write([]byte(out))
	})

	http.ListenAndServe(":6060", nil)
	
	// Save the catalog before exiting.
	Controller.EncodeCatalog()
}
