package main
import(
  "fmt"
  "net/http"
  "html/template"
)

// function name as key
var funcMap = template.FuncMap{
  // body as value
  "unescape": func (s string) template.HTML {
    return template.HTML(s)
  }, 
  "avg": func(n ...int) int {
    var total = 0
    for _, each := range n {
      total += each 
    }
    return total/len(n)
  },
}

func main()  {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
    // make a new instance template
    var tmpl = template.Must(template.New("view.html").
          // regist custom funtion to template
          Funcs(funcMap).
          // parsing template html
          ParseFiles("view.html"))
    if err := tmpl.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  fmt.Println("Server starting")
  http.ListenAndServe(":9000", nil)
}
