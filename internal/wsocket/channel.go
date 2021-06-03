package wsocket

import (
	"fmt"
	"log"

	"github.com/lorezi/golang-websocket/internal/dto"
)

var wsChan = make(chan dto.WSPayload)
var clients = make(map[dto.WebSocketConnection]string)

func ListenForWS(conn *dto.WebSocketConnection, cl map[dto.WebSocketConnection]string) {
	clients = cl
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload dto.WSPayload
	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do something
			log.Println(err)
		}
		payload.Conn = *conn
		// send payload to channel
		wsChan <- payload
	}
}

func ListenToWSChannel() {
	var res dto.WSJsonResponse

	for {
		e := <-wsChan
		res.Action = "Got here"
		res.Message = fmt.Sprintf("Some message... and action was %s", e.Action)
		broadcastToAll(res)
	}
}

func broadcastToAll(res dto.WSJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(res)
		if err != nil {
			log.Printf("websocket err %s", err)
			err = client.Close()
			if err != nil {
				log.Printf("error while closeing websocket %s", err)
			}
			delete(clients, client)
		}
	}
}
