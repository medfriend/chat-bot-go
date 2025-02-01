package zmq

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/medfriend/shared-commons-go/util/consul"
	"github.com/pebbe/zmq4"
)

func ConnZMQ(consulClient *api.Client) *zmq4.Socket {

	serviceInfo, _ := consul.GetKeyValue(consulClient, "CHATBOT")

	var resultServiceInfo map[string]string

	err := json.Unmarshal([]byte(serviceInfo), &resultServiceInfo)
	zmqPort := resultServiceInfo["SERVICE_PORT"]

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

	return socket
}
