package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

const serverAddr = ":8080"
const publicFolder = "./public"

func main() {
	log.SetFlags(log.Lshortfile)

	// starting server
	handler := http.FileServer(http.Dir(publicFolder))
	server := &http.Server{Addr: serverAddr, Handler: handler}
	go func() {
		log.Printf("listening on %v...", serverAddr)
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
