package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//create a log file
	logFile, err := os.OpenFile("out.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// starting server
	fmt.Println("Create listener...")
	defer fmt.Println("Stop listener")
	tcpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	var conCounter int = 0

	for conCounter < 5 {
		tcpCon, err := tcpListener.Accept()
		if err != nil {
			panic(err)
		}

		go cnHandler(tcpCon, conCounter)
		conCounter += 1
	}
}

func cnHandler(con net.Conn, id int) {
	defer con.Close()
	userName := fmt.Sprintf("user_%d", id)
	//log connection info
	log.Println("User connected: ", userName, " localAddr: ", con.LocalAddr().String(), " remoteAddr: ", con.RemoteAddr().String())

	//greatings
	fmt.Fprintln(con, "Hello user ", userName, " say something\n")
	fmt.Fprintln(con, "Print 'bye' for close connection\n")

	//handle response
	response := bufio.NewScanner(con)
	for response.Scan() {
		inLine := response.Text()
		if inLine == "bye" {
			log.Println("User ", userName, " close session")
			break
		}

		log.Println(userName, ":", inLine)
		fmt.Fprintf(con, "logging ok\n")
	}

	fmt.Println("Bye user ", userName)
}
