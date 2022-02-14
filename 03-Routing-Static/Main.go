package main
import(
	"fmt"
	"net/http"
)

func main(){
	http.Handle("/static/",
		// membungkus actual handler, menghapus prefix dari endpoint request
        // contoh tanpa prefix : /assets/static/site.css
        // contoh setelah prefix : /assets/site.css
		http.StripPrefix("/static/",
			// handler asli
			http.FileServer(http.Dir("assets"))))

        
        // learn method (err example)
        // http.Handle("/",http.FileServer(http.Dir("assets")))
        // http.Handle("/static",http.FileServer(http.Dir("assets")))

    fmt.Println("server started :9000")
	http.ListenAndServe(":9000", nil)
}
