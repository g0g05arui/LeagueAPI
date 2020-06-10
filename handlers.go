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

	NEVERMIND ??
	Problema era cu grupurile , dar cred ca aia se face in cadrul structurii de grup...
*/
func readHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	querry := p.ByName("key")
	if querry == "user" { ///By UID
		qID, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
		user, err := getUser(qID)
		var printData []byte
		if err == nil {
			printData, _ = json.Marshal(user)
		} else {
			printData, _ = json.Marshal(HTTPResponse{Response: "Unexistend user", Code: 404})
		}
		fmt.Fprintln(w, string(printData))
	} else if querry == "post" {
		fmt.Fprintf(w, "post")
	}
}

func addHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	querry := p.ByName("key")
	if querry == "user" {
		decoder := json.NewDecoder(r.Body)
		newU := User{0, "", "", "", "", 0, "", "", "", ""}
		decoder.Decode(&newU)
		response := addUser(newU)
		printData, _ := json.Marshal(response)
		fmt.Fprintln(w, string(printData))
	} else if querry == "post" {

	}
}
