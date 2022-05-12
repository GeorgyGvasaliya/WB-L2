package config

import (
	"flag"
	"log"
	"sync"
)

type Config struct {
	A         int
	B         int
	C         int
	RegString string
	Filename  string
	Count     bool
	I         bool
	V         bool
	F         bool
	N         bool
}

var once sync.Once
var instance Config

func NewConfig() Config {
	once.Do(func() {
		flag.IntVar(&instance.A, "A", 0, "печатать +N строк после совпадения")
		flag.IntVar(&instance.B, "B", 0, "печатать +N строк до совпадения")
		flag.IntVar(&instance.C, "C", 0, "печатать ±N строк вокруг совпадения")
		flag.BoolVar(&instance.Count, "c", false, "количество строк")
		flag.BoolVar(&instance.I, "i", false, "игнорировать регистр")
		flag.BoolVar(&instance.V, "v", false, "вместо совпадения, исключать")
		flag.BoolVar(&instance.F, "F", false, "точное совпадение со строкой, не паттерн")
		flag.BoolVar(&instance.N, "n", false, "печатать номер строки")
		flag.Parse()

		args := flag.Args()
		if len(args) == 2 {
			instance.RegString = args[0]
			instance.Filename = args[1]
			return
		}
		log.Fatalln("flags -> regexp string -> filename")
	})
	return instance
}
