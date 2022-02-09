package main
import(
  "fmt"
  "net/http"
  "html/template"
)

type Superhero struct {
  Name  string
  Alias string
  Friends []string
}

func (s Superhero) SayHello(from string, message string) string{
  return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
    var person = Superhero{
      Name:     "Akiekano",
      Alias:    "Hello zero two",
      Friends:  []string{"Moona", "Oolie", "Fubuki"},
    }

    var tmpl = template.Must(template.ParseFiles("view.html"))
    if err := tmpl.Execute(w, person); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  })

  fmt.Println("Started")
  http.ListenAndServe(":9000",nil)
}
