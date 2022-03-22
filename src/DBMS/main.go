package main

import (
	"GoDBMS/Controller"
	"net/http"
	"io/ioutil"
)

func main() {

	Controller.InitializeCatalog()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		body, _ := ioutil.ReadAll(r.Body)
		out := Controller.StartDBMS(string(body))
		w.Write([]byte(out))
	})

	http.ListenAndServe(":6060", nil)
}