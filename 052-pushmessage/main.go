package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var err error

func main() {
	http.Handle("/", http.HandlerFunc(indexHandler))
	http.Handle("/ws", http.HandlerFunc(wsHandler))
	http.ListenAndServe(":12345", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	go echo(c)
}

func echo(c *websocket.Conn) {
	for {
		if err = c.WriteJSON("Hellow world"); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("helloworld.html")
	t.Execute(w, "Hello World~~!!!")
}
