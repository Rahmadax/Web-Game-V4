package main

import (
	"fmt"
	inv "github.com/Rahmadax/Web-Game-V4/pkg/gameplay/inventory"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homepage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Homepage")
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Print("connection established...")

	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("we", wsEndpoint)
}

func main() {

}

func runTest() {

}

func createInventory() inv.Inventory {
	emptySlot := inv.ItemSlot{}
	//fullSlot0 := inv.ItemSlot{
	//	ItemId: "0",
	//	Amount: 1,
	//}

	fullSlot1 := inv.ItemSlot{
		ItemId: "1",
		Amount: 3,
	}

	currencies := inv.Currencies{Gold: 500}

	inventory := inv.Inventory{
		MaxSize:    8,
		Currencies: currencies,
		ItemSlots:  nil,
	}
	inventory.ItemSlots = append(inventory.ItemSlots, fullSlot1)
	inventory.ItemSlots = append(inventory.ItemSlots, emptySlot)
	inventory.ItemSlots = append(inventory.ItemSlots, emptySlot)
	inventory.ItemSlots = append(inventory.ItemSlots, emptySlot)
	inventory.ItemSlots = append(inventory.ItemSlots, emptySlot)
	inventory.ItemSlots = append(inventory.ItemSlots, fullSlot1)
	inventory.ItemSlots = append(inventory.ItemSlots, emptySlot)
	inventory.ItemSlots = append(inventory.ItemSlots, emptySlot)

	return inventory
}

//	fmt.Println("Starting Server")
//	setupRoutes()
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
