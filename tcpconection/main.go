package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//create log file
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Can't create log file: ", err)
		os.Exit(-1)
	}
	defer logFile.Close()

	//setup logging
	log.SetOutput(logFile)

	//starting tcp server
	log.Println("starting tcp server...")
	defer log.Println("shoutdown tcp server")

	tcpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer tcpListener.Close()

	log.Println("Wait for connection")

	tcpConnection, err := tcpListener.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	lAddr := tcpConnection.LocalAddr()
	rAddr := tcpConnection.RemoteAddr()

	fmt.Fprintln(tcpConnection, "GERARA HERE!!! YOUR FUCKING IDIOT! I KNOW U ADDR: ", lAddr.String(), rAddr.String())

	tcpConnection.Close()
	log.Println("Tcp session complete")
}
