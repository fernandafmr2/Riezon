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

	// respnse header
	w.Header().Set("Content-Type", "application/json")
	
	// encoder
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
}