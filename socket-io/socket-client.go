package main

import (
	"fmt"
	"log"
	"os"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func main() {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	uri := fmt.Sprintf("http://%s:%s", os.Getenv("URL"), os.Getenv("PORT"))
	log.Println(uri)

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("connection", func() {
		// lib error? this line not called
		log.Printf("on connect\n")
	})

	client.On("error", func() {
		log.Printf("on error\n")
	})

	client.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})

	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})
}
