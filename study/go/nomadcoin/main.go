package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RokiLord/nomadcoin/utils"
)

const port string = ":4000"

type URLDescription struct {
	URL         string
	Method      string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
	}
	b, err := json.Marshal(data)
	utils.HandleFunc(err)
	fmt.Printf("%s", b)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s", port)
	http.ListenAndServe(port, nil)

}
