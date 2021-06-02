package dto

// WSJsonResponse defines the response sent back from websocket
type WSJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}
