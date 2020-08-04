package app

import (
	"fmt"
	"log"
	"net/http"
)

func wsRoute(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("ERROR ReadMessage", err.Error())
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Println("ERROR WriteMessage", err.Error())
			return
		}
	}
}
