package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":8081"

	fmt.Println("server started at localhost:8081")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {

	// check valid basic auth or not
	if !Auth(w, r) {
		return
	}

	// check only GET request
	if !AllowOnlyGET(w, r) {
		return
	}

	// check request have parameter "id"
	if id := r.URL.Query().Get("id"); id != "" {
		// only the user with the desired id is used
		// as the return value
		OutputJSON(w, SelectStudent(id))
		return
	}

	// end point return all data user have
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

