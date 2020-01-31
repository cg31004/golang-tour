package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

//Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

//Message is define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	//Create a simple file server
	fs := http.FileServer(http.Dir("Public"))
	fmt.Println(fs)
	http.Handle("/", fs)

	//Cibfigure websocket route
	http.HandleFunc("/ws", handleConnections)

	//Start listening for incoming chat messages
	go handleMessages()

	//Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	//Upgraade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	//Make sure we close the connection when the function returns
	defer ws.Close()
	//Reguster our new client
	clients[ws] = true

	for {
		var msg Message
		//Read in  a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		//Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		//Grab the next message from the broadcast channel
		msg := <-broadcast
		//Send it out to every clent that is currently connnected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
