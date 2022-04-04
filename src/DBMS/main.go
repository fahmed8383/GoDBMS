package main

import (
	"GoDBMS/Controller"
	"GoDBMS/StorageLock"
	"net/http"
	"io/ioutil"
)

func main() {

	Controller.InitializeCatalog()
	StorageLock.InitializeLocks()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		body, _ := ioutil.ReadAll(r.Body)
		out := Controller.StartDBMS(string(body))
		w.Write([]byte(out))
	})

	http.ListenAndServe(":6060", nil)
}