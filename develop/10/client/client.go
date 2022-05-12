package myclient

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
)

func ReadRoutine(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	//функция завершения
	defer cancel()
	scanner := bufio.NewScanner(conn)

OUTER:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("break read")
			break OUTER
		default:
			if scanner.Err() != nil {
				fmt.Println(scanner.Err())
			}

			if !scanner.Scan() {
				log.Printf("CANNOT SCAN from conn")
				break OUTER
			}

			text := scanner.Text()
			log.Printf("from server: %s", text)
		}
	}

	log.Printf("finished read")
}

func WriteRoutine(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	defer cancel()
	scanner := bufio.NewScanner(os.Stdin)

OUTER:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("break write")
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			_, err := conn.Write([]byte(fmt.Sprintf("%s\n", str)))
			if err != nil {
				log.Printf("error on write to server: %v\n", err)
				break OUTER
			}
		}
	}

	log.Printf("finished write")
}
