package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tianjin University!\n"))
}

func main() {
	flag.Parse()
	router := mux.NewRouter()
	router.HandleFunc("/my", MyHandler)
	//http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8888", router))
}
