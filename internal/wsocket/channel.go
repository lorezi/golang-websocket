package wsocket

import (
	"fmt"
	"log"
	"sort"

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
		// res.Action = "Got here"
		// res.Message = fmt.Sprintf("Some message... and action was %s", e.Action)
		// broadcastToAll(res)

		switch e.Action {
		case "username":
			// get a list of all users and send it back via broadcast
			clients[e.Conn] = e.Username
			users := getUsers()
			res.Action = "list_users"
			res.ConnectedUsers = users
			broadcastToAll(res)
		}

	}
}

// getUsers builds users list and returns sorted users

func getUsers() []string {
	users := []string{}
	for _, v := range clients {
		users = append(users, v)
	}

	// sort the users
	sort.Strings(users)

	return users
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
