package awesomeProject

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the homepage")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

func main() {

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
			n, err := fmt.Fprintf(w, "Go Web App gestartet")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
		})
	*/

	http.HandleFunc("/", Home)
	http.HandleFunc("/", About)

	_ = http.ListenAndServe(":9000", nil)
}
