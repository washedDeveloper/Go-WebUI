package dependencies

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func IPC() {
	addr := flag.String("addr", "localhost:1234", "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	fmt.Println(u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Print(err)
		return
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(message)
	}
}
