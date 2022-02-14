package main
import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case "POST":
			w.Write([]byte("post"))
		case "GET":
			w.Write([]byte("get"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("Server starting")
	http.ListenAndServe(":9000", nil)
}