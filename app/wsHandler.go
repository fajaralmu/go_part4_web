package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

func sendBroadcastMessage(msg string) {
	broadcast <- msg
}

func sendMessageToClient(w http.ResponseWriter, r *http.Request, message string) {
	conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	if nil != err {
		log.Println("err: ", err.Error())
		return
	}
	for {
		// // Read message from browser
		// msgType, msg, err := conn.ReadMessage()
		// if err != nil {
		// 	log.Println("ERROR ReadMessage", err.Error())
		// 	return
		// }

		// Print the message to the console
		// fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		msg := []byte(message)
		fmt.Println("[]byte(message): ", msg)
		if err := conn.WriteMessage(1, msg); err != nil {
			log.Println("ERROR WriteMessage", err.Error())
			return
		}
	}
}

func wsRoute(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	clients[conn] = true
	defer conn.Close()
	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("ERROR ReadMessage", err.Error())
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s %v \n", conn.RemoteAddr(), string(msg), msgType)

		// // Write message back to browser
		// if err = conn.WriteMessage(msgType, msg); err != nil {
		// 	log.Println("ERROR WriteMessage", err.Error())
		// 	return
		// }
		sendBroadcastMessage(string(msg))
		log.Println("...")
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		log.Println("WILL handleMessages: ", msg)
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(struct{ Msg string }{msg})
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
