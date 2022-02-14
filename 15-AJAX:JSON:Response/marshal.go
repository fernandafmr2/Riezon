package main
import(
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("Starting...")
	http.ListenAndServe(":9000", nil)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := [] struct {
		Name 	string
		Age 	int
	} {
		{"Majitenshi towa", 17},
		{"Moona", 18},
		{"Akie", 22},
		{"Rushia", 18},
	}

	// convert data into json(data)
	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	// respnse header
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}