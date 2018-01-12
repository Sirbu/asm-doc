package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func opcode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	file := params["opcode"]

	if file != "style.css" {
		file += ".html"
	}

	fileResponse, err := os.Open("html/" + file)
	if err != nil {
		fmt.Fprintln(w, "Sorry, this opcode is not known. Check the spelling maybe !")

	} else {
		buffer, err := ioutil.ReadAll(fileResponse)
		if err != nil {
			fmt.Println(err.Error())
		}
		w.Write(buffer)
	}
}

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/x86/{opcode}", opcode)

	err := http.ListenAndServe("localhost:8000", mux)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("over and out")

}
