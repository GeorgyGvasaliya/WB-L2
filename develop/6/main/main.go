package main

import (
	"L2/develop/6/config"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	cfg := config.NewConfig()

	if cfg.Fields == 0 {
		log.Fatalln("-f must be > 0")
	}

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	res := Cut(text, cfg)
	fmt.Println(res)
}

func Cut(str string, cfg config.Config) string {
	if cfg.Separated {
		if !strings.Contains(str, cfg.Delim) {
			return ""
		}
	}
	spl := strings.Split(str, cfg.Delim)
	if cfg.Fields <= len(spl) {
		return spl[cfg.Fields-1] + "\n"
	}
	return ""
}
