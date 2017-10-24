package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

// RunHost does the cool hosty stuff
func RunHost(ip string) {
	ipAndPort := ip + ":" + port

	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Oops someone made a mistake", listenErr)
	}
	fmt.Println("Listening on", ipAndPort)

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Oops someone made a mistake", acceptErr)
	}
	fmt.Println("New connection accepted")

	for {
		handleHost(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Oops someone made a mistake", readErr)
	}
	fmt.Println("Messaaage:", message)

	fmt.Print("Send a cool messaaage: ")
	replyReader := bufio.NewReader(os.Stdin)
	message, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Oops someone made a mistake", replyErr)
	}
	fmt.Fprint(conn, message)
}

// RunGuest does the cool guesty stuff
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port

	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Oops someone made a mistake", dialErr)
	}

	for {
		handleGuest(conn)
	}
}

func handleGuest(conn net.Conn) {
	fmt.Print("Send a cool messaaage: ")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Oops someone made a mistake", readErr)
	}
	fmt.Fprint(conn, message)

	replyReader := bufio.NewReader(conn)
	replyMessage, replyError := replyReader.ReadString('\n')
	if replyError != nil {
		log.Fatal("Oops someone made a mistake", readErr)
	}
	fmt.Println("Messaaage:", replyMessage)
}
