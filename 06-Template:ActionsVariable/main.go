package main
import(
  "fmt"
  "net/http"
  "html/template"
)

type Info struct {
  Affiliation   string
  Address        string 
}

type Person struct {
  Name          string
  Gender        string
  Hobbies       []string
  Info          Info
}

func (t Info) GetAffiliationDetailInfo() string {
  return "have 31 waifu's"
} 

func main() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
    var person = Person{
      Name : "Akiekano",
      Gender : "Male",
      Hobbies: []string{"Reading", "coding", "learning"},
      Info: Info{"Moekano","Osaka Japan"},
    }

    var tmpl = template.Must(template.ParseFiles("view.html"))
    if err := tmpl.Execute(w, person); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    } 
  })

  //disable annoying warning
  // isTrue := true
  // _ = isTrue

  fmt.Println("Started")
  http.ListenAndServe(":9000",nil)
}
