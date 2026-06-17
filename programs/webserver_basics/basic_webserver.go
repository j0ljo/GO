package main 

import (
	"fmt"
	"net/http"
	
)

func main() {
	// define router 
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", home)
	
	fmt.Println("Server staring on :8080...")


	http.ListenAndServe(":8080", mux)
}

// define our 'home' handler function
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the page")
}


