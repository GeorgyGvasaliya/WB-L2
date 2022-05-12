package main

import (
	"L2/develop/5/config"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func getFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func main() {
	cfg := config.NewConfig()
	text, err := getFile(cfg.Filename)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(task(text, cfg))
}

func task(text []string, cfg config.Config) []string {
	var res []string
	pref := ""
	post := ""
	//ignore case
	if cfg.I {
		pref = "(?i)"
	}
	//current string not a pattern
	if cfg.F {
		pref += "^"
		post = "$"
	}
	reg, err := regexp.Compile(pref + cfg.RegString + post)
	if err != nil {
		log.Fatalln("error compile regular")
	}
	// after flag
	if cfg.A != 0 {
		for i, s := range text {
			if reg.MatchString(s) {
				if cfg.A <= len(text)-i {
					res = append(res, fmt.Sprintf("A: %s", text[i:i+cfg.A+1]))
					break
				}
				res = append(res, fmt.Sprintf("A: %s", text[i:]))
				break
			}
		}
	}
	// before flag
	if cfg.B != 0 {
		for i, s := range text {
			if reg.MatchString(s) {
				if cfg.B-1 < i {
					res = append(res, fmt.Sprintf("B: %s", text[i-cfg.B:i+1]))
					break
				}
				res = append(res, fmt.Sprintf("B: %s", text[:i+1]))
				break
			}
		}
	}
	// after + before flag
	if cfg.C != 0 {
		for i, s := range text {
			if reg.MatchString(s) {
				start := 0
				end := len(text)
				if cfg.C-1 < i {
					start = i - cfg.C
				}
				if cfg.C <= end-i {
					end = i + cfg.C + 1
				}

				res = append(res, fmt.Sprintf("C: %s", text[start:end]))
				break
			}
		}
	}
	// count match flag
	if cfg.Count {
		cnt := 0
		for _, s := range text {
			cnt += len(reg.FindAllString(s, -1))
		}
		if cfg.V {
			res = append(res, fmt.Sprintf("Count -c -v (invert):%d", len(text)-cnt))
		} else {
			res = append(res, fmt.Sprintf("Count -c:%d", cnt))
		}
	}
	// num of matches (+invert num) flag
	if cfg.N {
		var n []int
		for i, s := range text {
			if reg.MatchString(s) {
				n = append(n, i+1)
			}
		}
		res = append(res, fmt.Sprintf("Match nunbers -n:%d", n))
		cfg.N = false
	}
	return res
}
