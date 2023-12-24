package main

import (
	"bufio"
	"io"
	"log"
	"net"

	"github.com/Antonov-guap/auth-pow/pkg/pow"
)

const serverAddr = "localhost:8080"

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("error connecting to server:", err)
	}
	defer func() {
		_ = conn.Close()
		log.Println("connection closed")
	}()
	log.Println("connection opened to", serverAddr)

	err = pow.Read(conn, func(r io.Reader) error {
		reader := bufio.NewReader(r)
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read response from server: %v", err)
		}
		log.Print("server response: ", response)
		return nil
	})
}
