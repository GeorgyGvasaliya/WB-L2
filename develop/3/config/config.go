package config

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

type Config struct {
	K        int
	Filename string
	N        bool
	R        bool
	U        bool
}

var once sync.Once
var instance Config

func NewConfig() Config {
	once.Do(func() {
		flag.IntVar(&instance.K, "k", 0, "указание колонки для сортировки")
		flag.BoolVar(&instance.N, "n", false, "сортировать по числовому значению")
		flag.BoolVar(&instance.R, "r", false, "сортировать в обратном порядке")
		flag.BoolVar(&instance.U, "u", false, "не выводить повторяющиеся строки")
		flag.Parse()

		args := flag.Args()
		if len(args) != 1 {
			usage()
		}
		instance.Filename = args[0]
	})
	return instance
}

func usage() {
	fmt.Println("Usage: [-k указание колонки для сортировки] [-n сортировать по числовому значению] " +
		"[-r сортировать в обратном порядке] [-u не выводить повторяющиеся строки]")
	os.Exit(0)
}
