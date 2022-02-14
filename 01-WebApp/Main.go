package main
import "fmt"
import "net/http"

// rules for make handler
func handlerIndex(w http.ResponseWriter, r *http.Request){
	var message = "Welcome"
	w.Write([]byte(message))

}

func handlerHello(w http.ResponseWriter, r *http.Request){
	var message = "Hello World!"
	w.Write([]byte(message))
}

func main(){

	// routing(route, handler)
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)

	// start with server
	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}
