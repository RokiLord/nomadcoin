package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

//response writer = data that wants to send to users
func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from Home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	//error handling
	log.Fatal(http.ListenAndServe(port, nil))
}
