package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadFile("html/list.html")
		w.Write(data)
	})
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./js/"))))
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./js/"))))
	log.Fatal(http.ListenAndServe(":8000", router))
}
