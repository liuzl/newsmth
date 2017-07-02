package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	start, _ := strconv.ParseUint(r.FormValue("start"), 10, 32)
	limit, err := strconv.ParseUint(r.FormValue("limit"), 10, 32)
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	rv := struct {
		Start uint64 `json:"start"`
		Limit uint64 `json:"limit"`
		Error string `json:"error"`
	}{
		Start: start,
		Limit: limit,
		Error: errStr,
	}
	mustEncode(w, rv)
}

func main() {
	flag.Parse()
	router := mux.NewRouter()
	router.HandleFunc("/echo", MyHandler)
	//http.Handle("/", router)
	glog.Fatal(http.ListenAndServe(":8888", router))
}
