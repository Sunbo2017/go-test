package main

import (
	_ "embed"
	"github.com/lxzan/gws"
	"log"
	"net/http"
	"time"
)

func main() {
	upgrader := gws.NewUpgrader(&Handler{}, &gws.ServerOption{
		CompressEnabled:  true,
		CheckUtf8Enabled: true,
		Recovery:         gws.Recovery,
	})
	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(writer, request)
		if err != nil {
			return
		}
		go func() {
			socket.ReadLoop()
		}()
		for {
			log.Println("start hello----")
			socket.WriteMessage(gws.OpcodeText, []byte("hello world"))
			time.Sleep(10 * time.Second)
		}

	})
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

type Handler struct {
	gws.BuiltinEventHandler
}

func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	log.Println("on ping,write pong====")
	_ = socket.WritePong(payload)
}

func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	//_ = socket.WriteMessage(message.Opcode, message.Bytes())
	log.Printf("msg===%+v", string(message.Bytes()))
}
