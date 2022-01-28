package awesomeProject

import (
	"fmt"
	"net/http"
)

const portNumber = ":9000"

// Home returns the Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the homepage")
}

// About returns the About Page
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/", About)

	fmt.Println("Starting App on port 9000...")
	_ = http.ListenAndServe(portNumber, nil)

}
