package main

import (
	"chat-bot-go/config"
	"chat-bot-go/openAI"
	"encoding/json"
	"fmt"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/medfriend/shared-commons-go/util/env"
	"github.com/pebbe/zmq4"
)

func main() {
	// Crear un socket PULL

	env.LoadEnv()

	consulClient := consul.ConnectToConsulKey("CHATBOT")
	serviceInfo, _ := consul.GetKeyValue(consulClient, "CHATBOT")

	var resultServiceInfo map[string]string

	err := json.Unmarshal([]byte(serviceInfo), &resultServiceInfo)
	zmqPort := resultServiceInfo["SERVICE_PORT"]
	config.Init(resultServiceInfo["KEY"])

	socket, err := zmq4.NewSocket(zmq4.PULL)
	if err != nil {
		panic(err)
	}
	defer socket.Close()

	zmqConn := fmt.Sprintf("tcp://localhost:%s", zmqPort)

	// Conectar al servidor PUSH
	err = socket.Connect(zmqConn)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Connected to %s", zmqConn))

	// Escuchar mensajes

	for {
		msg, err := socket.Recv(0)
		if err != nil {
			fmt.Printf("Error recibiendo mensaje: %v\n", err)
			continue
		}
		openAI.MakeRequest(msg)
	}
}
