package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/RokiLord/nomadcoin/blockchain"
)

//port
const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

//response writer = data that wants to send to users
func home(rw http.ResponseWriter, r *http.Request) {
	//parsing home.html to tmpl

	//get blocks data
	data := homeData{"Home", blockchain.GetBlockChain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*gohtml"))
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	//error handling
	log.Fatal(http.ListenAndServe(port, nil))
}
