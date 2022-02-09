package main
import (
	"fmt"
	"net/http"
)

func main(){
	// closure
	handlerIndex := func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("hello"))
	}

	// anonymous func
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello again"))
	})


	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)

	fmt.Println("server started at localhost:9000")
	// err := http.ListenAndServe(":9002", nil)

	// connect 
	server := new(http.Server)
	server.Addr = ":9000"
	err := server.ListenAndServe()
	if err != nil{
		fmt.Println(err.Error())
	}
}

// dalam routing handler bisa beru\pa gungsi, closure, atau\pun anonymous func