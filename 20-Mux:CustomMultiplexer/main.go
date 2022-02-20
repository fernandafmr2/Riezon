package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := new(CustomMux)

	mux.HandleFunc("/student", ActionStudent)

	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	server := new(http.Server)
	server.Addr = ":8081"
	server.Handler = mux

	fmt.Println("server started at localhost:8081")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())


}

func OutputJSON(w http.ResponseWriter, o interface{}) {

	// convertion data object or slice to JSON
	res, err := json.Marshal(o)
	
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

