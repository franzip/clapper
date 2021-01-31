package hub

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

const (
	CLIENT_JOINED = "CLIENT_JOINED"
	CLIENT_LEFT   = "CLIENT_LEFT"
)

type Client struct {
	Id         string
	Username   string
	Connection *websocket.Conn
}

type Hub struct {
	clients map[string]Client
}

type Message struct {
	Topic string `json:"topic"`
	User  string `json:"user"`
}

func (client Client) send(message []byte) error {
	return client.Connection.WriteMessage(1, message)
}

func Init() Hub {
	var hub = Hub{clients: make(map[string]Client)}
	return hub
}

func (hub *Hub) ProcessMessage(messageData []byte) error {
	var message Message

	err := json.Unmarshal(messageData, &message)

	if err != nil {
		return err
	}

	fmt.Println(message)
	switch message.Topic {
	case CLIENT_LEFT:
		hub.RemoveClient(message.User)
	}

	return nil
}

func (hub *Hub) RemoveClient(username string) {
	delete(hub.clients, username)
}

func (hub *Hub) AddClient(client Client) error {
	hub.clients[client.Username] = client

	message := Message{Topic: CLIENT_JOINED, User: client.Username}

	return hub.broadcast(message)
}

func (hub *Hub) broadcast(message Message) error {
	jsonMessage, err := json.Marshal(message)

	if err != nil {
		return err
	}
	fmt.Println(len(hub.clients))
	for _, client := range hub.clients {
		err = client.send(jsonMessage)
		if err != nil {
			fmt.Println("error broadcasting: ", err)
			return err
		}
	}

	return nil
}
