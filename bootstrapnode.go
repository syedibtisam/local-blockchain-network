package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	SERVER_HOST    = "localhost"
	Bootstrap_PORT = "8000"
	SERVER_TYPE    = "tcp"
)

var port = 1000

func main() {

	fmt.Println("Bootstrap node running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+Bootstrap_PORT) // listening for clients
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + Bootstrap_PORT)
	fmt.Println("Waiting for Nodes to send them their ports...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// fmt.Println("client connected")
		go processClient(connection)
	}
}
func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	if "get_total_ports" == string(buffer[:mLen]) { // if want to get update on neighbours
		con, err := connection.Write([]byte(string(strconv.Itoa(port))))
		if err != nil {
			fmt.Println(con)
		}
		fmt.Println("node updated with neighbours")
		fmt.Println("client connected, updated with total neighbours and disconnected")
		connection.Close()
	} else if "get_port" == string(buffer[:mLen]) { // when new node connected to get the port on which he will be listening
		_, err := connection.Write([]byte(string(strconv.Itoa(port))))
		if err != nil {
			fmt.Println(err)
		}
		port++
		// fmt.Println(con)
		fmt.Println("client connected, got port number" + strconv.Itoa(port-1) + " and disconnected")
		connection.Close()

	}
	// fmt.Println("Node received port:", port)
	// fmt.Println(", and serval neighbours")
	// ("Thanks! Got your message:" + string(buffer[:mLen]))
}
