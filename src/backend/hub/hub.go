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
	Id         string          `json:"id"`
	Username   string          `json:"username"`
	Connection *websocket.Conn `json:"-"`
}

type Hub struct {
	clients map[string]Client
}

type Message struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
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

	switch message.Topic {
	case CLIENT_LEFT:
		err := hub.RemoveClient(message.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (hub *Hub) RemoveClient(username string) error {
	delete(hub.clients, username)

	clients, err := json.Marshal(hub.getAllClients())

	if err != nil {
		fmt.Println(err)
		return err
	}

	message := Message{Topic: CLIENT_LEFT, Data: string(clients)}

	return hub.broadcast(message)
}

func (hub *Hub) AddClient(client Client) error {
	hub.clients[client.Username] = client

	clients, err := json.Marshal(hub.getAllClients())

	if err != nil {
		fmt.Println(err)
		return err
	}

	message := Message{Topic: CLIENT_JOINED, Data: string(clients)}

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

func (hub *Hub) getAllClients() []Client {
	var clients []Client

	for _, client := range hub.clients {
		clients = append(clients, client)
	}

	return clients
}
