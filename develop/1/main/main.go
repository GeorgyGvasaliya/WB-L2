package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	Mycurrtime()
}

func Mycurrtime() {

	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Printf("Точное время: %s / Текущее время: %s", t, time.Now())
}
