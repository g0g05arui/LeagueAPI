package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/*
	Problema cu securitate ??
	Oricine poate citi un user
	----TODO----
	Sa primesc in request si UID si SessionID al userului care face request-ul
	Verifica daca sunt compatibile UID si SID si dupa verific ce date poate sa primeasca UID-ul respectiv
	despre UID-ul din request ( grupuri / houses /etc...)
	----END------

*/
func readHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	querry := p.ByName("key")
	fmt.Println(querry)
	if querry == "user" { ///By UID
		qID, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
		printData, _ := json.Marshal(getUser(qID))
		fmt.Fprintln(w, string(printData))
	} else if querry == "post" {
		fmt.Fprintf(w, "post")
	}
}

/*
	----TODO----
	UID trebuie generate random
	----END-----
*/

func addHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	querry := p.ByName("key")
	fmt.Println(querry)
	if querry == "user" {
		decoder := json.NewDecoder(r.Body)
		newU := User{0, "", "", ""}
		decoder.Decode(&newU)
		response := addUser(newU)
		printData, _ := json.Marshal(response)
		fmt.Fprintln(w, string(printData))
		fmt.Println(newU)
	} else if querry == "post" {

	}
}
