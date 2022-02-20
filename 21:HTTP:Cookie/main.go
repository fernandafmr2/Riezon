package main

import "net/http"

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	html := "Hello World! "
	w.Write([]byte(html))
}

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("ithinkidroppedacookie")
	if err != nil {
		w.Write([]byte("error in reading cookie : " + err.Error() + "\n"))
	} else {
		value := c.Value
		w.Write([]byte("cookie has : " + value + "\n"))
	}
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "ithinkidroppedacookie",
		MaxAge: -1} //delete cookie now
		http.SetCookie(w, &c)

		w.Write([]byte("old cookie deleted!\n"))
}

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "ithinkidroppedacookie",
		Value:  "thedroppedcookiehasgoldinit",
		MaxAge: 3600}
		http.SetCookie(w, &c)

		w.Write([]byte("new cookie created!\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHelloWorld)
	mux.HandleFunc("/readcookie", ReadCookie)
	mux.HandleFunc("/deletecookie", DeleteCookie)
	mux.HandleFunc("/createcookie", CreateCookie)
	fmt.Println("starting..")
	http.ListenAndServe(":8080", mux)
}