package main
import(
	"fmt"
	"net/http"
	"html/template"
	"path/filepath"
	"os"
	"io"
)

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("Starting")
	http.ListenAndServe(":9000", nil)
}

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var tmpl = template.Must(template.ParseFiles("view.html"))
	var err = tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// ParseMultipartForm = parsing form data when including file
	// 1024 arg mean max memory
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get file
	alias := r.FormValue("alias")

	// get file upload
	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// make a new file
	filename := handler.Filename
	if alias != "" {
		// Ext for get file extension
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	// making path
	fileLocation := filepath.Join(dir, "files", filename)

	// os.OpenFile mean open file
	// WRONLY mean read and write only, and make if not found
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w , err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	// io.Copy fills the contents of the first parameter file
	// 		 with the contents of the second parameter
	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("done"))
}

