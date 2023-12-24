package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"

	"github.com/Antonov-guap/auth-pow/pkg/pow"
)

const (
	listenAddr = "localhost:8080" // Адрес и порт сервера
)

func main() {
	words, err := readWordsCollection()
	if err != nil {
		log.Fatal("unable to read words file:", err)
	}

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal("unable to listen for tcp:", err)
	}
	defer func() {
		_ = listener.Close()
		log.Println("server stopped")
	}()

	log.Println("server is listening to", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection:", err)
			continue
		}
		log.Println("incomming connection")

		go func() {
			defer func() {
				_ = conn.Close()
				log.Println("connection closed")
			}()

			err = pow.Handle(conn, words.writeRandomWordOfWisdom)
			if err != nil {
				log.Println("error handling connection:", err)
			}
		}()
	}
}

//go:embed quotes.txt
var quotesFile string

type wordsCollection []string

func readWordsCollection() (wordsCollection, error) {
	scanner := bufio.NewScanner(strings.NewReader(quotesFile))

	var collection wordsCollection
	for scanner.Scan() {
		collection = append(collection, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read embedded file: %w", err)
	}

	return collection, nil
}

func (c wordsCollection) writeRandomWordOfWisdom(w io.Writer) error {
	i := rand.Intn(len(c))
	randomWord := c[i]

	if _, err := fmt.Fprintln(w, randomWord); err != nil {
		return fmt.Errorf("failed to write random word to client: %w", err)
	}
	log.Println("random word sent to client:", randomWord)

	return nil
}
