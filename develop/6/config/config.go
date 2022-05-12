package config

import (
	"flag"
	"sync"
)

type Config struct {
	Fields    int
	Delim     string
	Separated bool
}

var once sync.Once
var instance Config

func NewConfig() Config {
	once.Do(func() {
		flag.IntVar(&instance.Fields, "f", 0, "fields")
		flag.StringVar(&instance.Delim, "d", "\t", "delim")
		flag.BoolVar(&instance.Separated, "s", false, "separated")
		flag.Parse()
	})
	return instance
}
