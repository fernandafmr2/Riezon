package main
import (
  "fmt"
  "net/http"
  "html/template"
  "path"
)

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
    // menggabung file dan folder menjadi sebuah path 
    var filepath = path.Join("views","index.html")
    // parsing file template 
    var tmpl, err = template.ParseFiles(filepath)
    
    if err != nil {
      // menandai ResponseWriter apabila terjadi eror 
      // StatusInternalServerError represent 500-Internal server eror 
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    // data template
    var data = map[string]interface{}{
      "title":"Learning Golang Web", 
      "nama":"Akiekano",
    }

    // menyisipkan data pada template dan menampilkan ke browser
    err = tmpl.Execute(w, data)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  // handler static file 
  http.Handle("/static/",
    http.StripPrefix("/static/",
      http.FileServer(http.Dir("assets"))))

  fmt.Println("Server started at localhost:9000")
  http.ListenAndServe(":9000", nil)
}
