package dto

import "github.com/gorilla/websocket"

// WSJsonResponse defines the response sent back from websocket
type WSJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

// WSPayloads defines the request sent to the websocket
type WSPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

// WebSocketConnection struct
type WebSocketConnection struct {
	*websocket.Conn
}
