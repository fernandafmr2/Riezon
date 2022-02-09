package main

import(
 "fmt"
 "net/http"
 "html/template"
)

// make alias
type M map[string]interface{}

func main()  {
  
  // parsing all files
  // var tmpl, err = template.ParseGlob("views/*")
  // 
  // if err != nil {
  //   panic(err.Error())
  //   return
  // }
  //
  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
  //   var data = M{"name":"Batman"}
  //   err = tmpl.ExecuteTemplate(w, "index", data)
  //   if err != nil {
  //     http.Error(w, err.Error(), http.StatusInternalServerError)
  //   }
  // })
  // http.HandleFunc("/about", func (w http.ResponseWriter, r* http.Request)  {
  // http.HandleFunc("/abt", func (w http.ResponseWriter, r* http.Request)  {
  //   var data = M{"name":"Batman"}
  //   err = tmpl.ExecuteTemplate(w , "about", data)
  //   if err != nil {
  //    http.Error(w, err.Error(), http.StatusInternalServerError)
  //   }
  // })

  // parsing seleted files
  http.HandleFunc("/index", func (w http.ResponseWriter, r *http.Request)  {
    var data = M{"name":"Batman"}
    // parsefiles into to detect eror when must execute
    var tmpl = template.Must(template.ParseFiles(
      "views/index.html",
      "views/_header.html",
      "views/_message.html",
    ))

    var err = tmpl.ExecuteTemplate(w, "index", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  http.HandleFunc("/about", func (w http.ResponseWriter, r *http.Request)  {
    var data = M{"name":"Batman"} 
    // parsefiles into to detect eror when must execute
    var tmpl = template.Must(template.ParseFiles(
      "views/about.html",
      "views/_header.html",
      "views/_message.html",
    ))

    var err = tmpl.ExecuteTemplate(w, "about", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })


  fmt.Println("Server started")
  http.ListenAndServe(":9000",nil)
}
