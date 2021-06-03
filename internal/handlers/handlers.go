package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lorezi/golang-websocket/internal/dto"
	"github.com/lorezi/golang-websocket/internal/utils"
	"github.com/lorezi/golang-websocket/internal/wsocket"
)

var upgradeConn = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Home displays the home page
func Home(w http.ResponseWriter, r *http.Request) {
	err := utils.RenderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}

}

// WSEndPoint upgrades connection to websocket
// The WSEndPoint is called from the home page
func WSEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConn.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client connected to the endpoint")

	res := dto.WSJsonResponse{
		Message: `<em><small>Connected to server</small><em>`,
	}

	conn := dto.WebSocketConnection{Conn: ws}

	clients := make(map[dto.WebSocketConnection]string)
	clients[conn] = ""

	err = ws.WriteJSON(res)
	if err != nil {
		log.Println(err)
	}

	// listen for websocket
	wsocket.ListenForWS(&conn, clients)
}
