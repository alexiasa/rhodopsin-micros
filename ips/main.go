package main

import (
	"log"
	"net/http"
	"github.com/urfave/negroni"
	"rhodopsin-micros/ips/routers"
)

func main() {

	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server { Handler: n}

	log.Println("Listening...")
	server.ListenAndServe()


}