package main

import (
	myclient "L2/develop/10/client"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	host    string
	port    int
	timeout time.Duration
)

//проверка наличия порта и хоста
func parseArgs(args []string) error {
	//если хост или порт не указанны
	if len(args) < 2 {
		return errors.New("should be host and port")
	}

	host := args[0]
	//если хост не указан
	if host == "" {
		return errors.New("host should be not empty")
	}
	//если порт не указан
	port, _ := strconv.Atoi(args[1])
	if port == 0 {
		return errors.New("port can't be zero")
	}

	return nil
}

func main() {
	//флаг timeout по умолчанию 10сек
	flagtime := flag.String("timeout", "10s", "should be timeout there?")

	flag.Parse()
	err := parseArgs(flag.Args())
	if err != nil {
		fmt.Println(err)
	}
	//задаем хост
	host = flag.Arg(0)
	//задаем номер порта
	port, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println("error name port transform")
	}

	timeout, err = time.ParseDuration(*flagtime)
	if err != nil {
		fmt.Println("error time  transform")
	}

	//инициализируем параметры
	addr := fmt.Sprintf("%s:%v", host, port)
	fmt.Printf("trying %s, timeout: %s...\n", addr, timeout)

	//установка timeout соединения
	dialer := &net.Dialer{Timeout: timeout}

	//пустой главный контекст not-nil
	ctx := context.Background()
	//производный контекст с новым done каналом, для завершения работы
	ctx, cancel := context.WithCancel(ctx)
	//создание соединения, с контекстом
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	fmt.Printf("connected to %s\n", addr)

	//чтение из соединения
	go myclient.ReadRoutine(ctx, conn, cancel)
	//запись в соединение
	go myclient.WriteRoutine(ctx, conn, cancel)

	//ждем сигнала завершения
	select {
	case <-ctx.Done():
		fmt.Println("shutdown signal received")
	}
	//закрыть соединение
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection closed")
}
