package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flaga := flag.Int("A", 0, "after")
	flagb := flag.Int("B", 0, "before")
	flagC := flag.Int("C", 0, "context")
	flagc := flag.Bool("c", false, "count")
	flagi := flag.Bool("i", false, "ignoreCase")
	flagv := flag.Bool("v", false, "invert")
	flagf := flag.Bool("F", false, "fixed")
	flagn := flag.Bool("n", false, "lineNum")

	flag.Parse()

	var strlist []string
	pos := -1
	filename := flag.Arg(0)
	lookFor := flag.Arg(1)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	//считывам построчно, добавляем в массив
	for scanner.Scan() {
		//если флаг i приводим  все к нижнему регистру
		if *flagi {
			strlist = append(strlist, strings.ToLower(scanner.Text()))
			lookFor = strings.ToLower(lookFor)
		} else {
			strlist = append(strlist, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//реализация флага f точное совпадение со строкой
	if *flagf {
		pos = FindFullStr(strlist, lookFor)
	} else {
		pos = FindStr(strlist, lookFor) //сначала поиск строки с совпадением
	}

	//реализация флага n печать номера строки
	if *flagn {
		if pos == -1 {
			fmt.Println("No such string")
		} else {
			fmt.Println("Number of string: ", pos)
		}
	}

	//реализация флага c печать количества совпадений
	if *flagc {
		count := FindCount(strlist, lookFor)
		fmt.Println("Count match: ", count)
	}

	//реализация флага v удаление совпадений
	if *flagv {
		strlist = RemoveMatch(strlist, lookFor)
		fmt.Println(strlist)
		return
	}

	//реализация флага a
	if *flaga != 0 {
		strlist = FindStringAfter(strlist, pos, *flaga)
	}

	//реализация флага b
	if *flagb != 0 {
		strlist = FindStringBefore(strlist, pos, *flagb)
	}

	//реализация флага C
	if *flagC != 0 {
		strlist = FindStringAround(strlist, pos, *flagC)

	}

	// печать резалта
	fmt.Println(strlist)
}

func FindFullStr(strlist []string, lookFor string) int {
	for i, curstr := range strlist {
		if curstr == lookFor {
			//вернуть индекс строки
			return i
		}
	}
	return -1
}

func FindStr(strlist []string, lookFor string) int {
	for i, curstr := range strlist {
		if strings.Contains(curstr, lookFor) == true {
			return i
		}
	}
	return -1
}

func FindCount(strlist []string, lookFor string) int {
	count := 0
	for _, curstr := range strlist {
		if strings.Contains(curstr, lookFor) == true {
			count++
		}
	}
	return count
}

func RemoveMatch(strlist []string, lookFor string) []string {
	newtext := strings.Join(strlist, " ")
	strlist = strings.Split(newtext, " ")
	newlist := []string{}

	for _, x := range strlist {
		if x != lookFor {
			newlist = append(newlist, x)
		}
	}
	return newlist
}

func FindStringAfter(strlist []string, pos int, N int) []string {
	changedList := []string{}
	//проверка на количество строк для флага a
	// подразумевается что введенное N валидно))
	for i := pos; i <= pos+N; i++ {
		changedList = append(changedList, strlist[i])
	}
	return changedList
}

func FindStringBefore(strlist []string, pos int, N int) []string {
	changedList := []string{}
	//проверка на количество строк для флага b
	for i := pos - N; i <= pos; i++ {
		changedList = append(changedList, strlist[i])
	}
	return changedList
}

func FindStringAround(strlist []string, pos int, N int) []string {
	changedList := []string{}
	for i := pos - N; i <= pos+N; i++ {
		changedList = append(changedList, strlist[i])
	}
	return changedList
}
