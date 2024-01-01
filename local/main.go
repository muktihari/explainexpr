package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	port := ":9090"
	fmt.Println("Running at: ", port)

	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt)
	go func() {
		log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir("../."))))
	}()
	<-q
}
