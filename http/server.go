package main

import (
	"io"
	log "github.com/Sirupsen/logrus"
	"net"
	"time"
	"os"
	"fmt"
)


func main() {

	// Resolves the address of the server port and listens
	addr, err := net.ResolveTCPAddr("tcp", "10.0.0.1:80")
	checkErr(err)
	listener, err := net.ListenTCP("tcp", addr)
	checkErr(err)
	createLog()
	log.Info("Server started.")

	// Indefinitely handle connections
	for {
		newConn, err := listener.AcceptTCP()
		checkErr(err)

		logClientJoined(newConn)
		checkErr(err)

		// Handle connection in a new goroutine
		go handleConn(newConn)
	}
}

// REQUIRES: conn is a pointer to a valid, open TCP connection
// MODIFIES: conn
// EFFECTS:	 Writes a sample message to the connection.
func handleConn(conn *net.TCPConn) {

	for {
		connIsClosed(conn)
		sampleMessage := []byte("Hello!\n")
		_, err := conn.Write(sampleMessage)
		checkErr(err)
		time.Sleep(1000 * time.Millisecond)
	}
}

// REQUIRES: conn is a pointer to a valid, open TCP connection
// EFFECTS:  Logs the event where a client joins.
func logClientJoined(conn *net.TCPConn) {
	fmt.Println("server.go: Client joined from %s", conn.RemoteAddr())
}

// EFFECTS:	 Handles any non-nil errors by printing them.
func checkErr(err error) {
	if err != nil {
		log.Error("Error: server.go: %s", err.Error())
	}
}

func connIsClosed(c *net.TCPConn) {
	c.SetReadDeadline(time.Now())
	var one []byte
	if _, err := c.Read(one); err == io.EOF {
		log.Info("Client disconnect: %s", c.RemoteAddr())
		c.Close()
		c = nil
	} else {
		var zero time.Time
		c.SetReadDeadline(zero)
	}
}

func createLog(){
	if f, err := os.OpenFile("/tmp/run_load.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err == nil {
		log.SetOutput(f)
	} else {
		fmt.Printf("Failed to open log file: %s",  err.Error())
		os.Exit(1)
	}
}