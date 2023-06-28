package main

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	io := socketio.NewServer(nil)


	//sockets
	io.OnConnect("/", func(so socketio.Conn) error {
		// so.SetContext("")
		// so.Join("chat_room")
		fmt.Println("New User Connected", so.ID())
		return nil
	})

	// io.OnEvent("/", "chat-message", func(so socketio.Conn, msg string){
	// 	fmt.Println(msg)
	// 	// serve.emit("chat-message")
	// 	io.BroadcastToRoom("/","chat_room", "chat-message", msg)

	// })

	go io.Serve()
	defer io.Close()

	//Modulo Http
	http.Handle("/socket.io/", io)
	http.Handle("/", http.FileServer(http.Dir("./angular-socket/dist/angular-socket")))
	log.Println("Server on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}