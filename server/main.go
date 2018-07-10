package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

const addr = ":8080"

func main() {
	// starting server
	h := http.FileServer(http.Dir("."))
	log.Printf("listening on %v...", addr)
	server := &http.Server{Addr: addr, Handler: h}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// listening for signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c

	// closing server
	log.Println("closing server...")
	server.Close()
}
