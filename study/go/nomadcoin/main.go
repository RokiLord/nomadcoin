package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/RokiLord/nomadcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

//response writer = data that wants to send to users
func home(rw http.ResponseWriter, r *http.Request) {
	//parsing home.html to tmpl
	tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	//get blocks data
	data := homeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	//error handling
	log.Fatal(http.ListenAndServe(port, nil))
}
