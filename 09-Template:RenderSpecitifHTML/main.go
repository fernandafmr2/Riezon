 package main 
 import(
   "fmt"
   "net/http"
   "html/template"
 )

 func main()  {
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
     var tmpl = template.Must(template.New("index").ParseFiles("view.html"))
     if err := tmpl.Execute(w, nil); err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
     }
   })

   http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
     var tmpl = template.Must(template.New("test").ParseFiles("view.html"))
     if err := tmpl.Execute(w, nil); err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
     }
   })

   fmt.Println("Starting")
   http.ListenAndServe(":9000",nil)
 }
