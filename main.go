package main

import (
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

var wg sync.WaitGroup

/*
	@Note Comunicatia nu este inca secure pentru ca folosim http
	----TODO----
	cumparat un domeniu + generat certificat SSL ca sa putem folosi https
	ca sa fie cryptate requesturi le
	----END-----
*/
func main() {
	serverCRUD := httprouter.New()
	serverCRUD.GET("/read/:key", readHandler)
	serverCRUD.POST("/add/:key", addHandler)
	wg.Add(1)
	go func() {
		http.ListenAndServe(":8555", serverCRUD)
		wg.Done()
	}()
	wg.Wait()
}
