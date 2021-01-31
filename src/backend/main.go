package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"

	"github.com/franzip/clapper/hub"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{}
var messageHub = hub.Init()

func getUuid() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	client := hub.Client{Id: getUuid(), Connection: c, Username: user}

	err = messageHub.AddClient(client)

	if err != nil {
		log.Print("error adding client:", err)
		return
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		err = messageHub.ProcessMessage(message)
		if err != nil {
			log.Println("error processing message:", err)
			break
		}
	}
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	entryPoint := filepath.Join(currentDir, "..", "..", "public", "index.html")
	http.ServeFile(w, r, entryPoint)
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	publicAssetsDir := filepath.Join(currentDir, "..", "..", "public")
	buildAssetsDir := filepath.Join(currentDir, "..", "..", "build")
	publicAssetFileServer := http.FileServer(http.Dir(publicAssetsDir))
	buildAssetsDirFileServer := http.FileServer(http.Dir(buildAssetsDir))
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", appHandler)
	http.Handle("/public/", http.StripPrefix(strings.TrimRight("/public/", "/"), publicAssetFileServer))
	http.Handle("/build/", http.StripPrefix(strings.TrimRight("/build/", "/"), buildAssetsDirFileServer))
	log.Fatal(http.ListenAndServe(*addr, nil))
	fmt.Println("Booting server...")
}
