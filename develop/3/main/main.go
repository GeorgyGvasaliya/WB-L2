package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	k := flag.Int("k", 1, "колонка для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "обратная сортировка")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	filename := flag.Arg(0)

	//  read data and transform it to matrix
	data := PrepareData(filename)

	if *u {
		data = DeleteDuplicates(data)
	}

	data = Sort(data, *k, *n, *r)

	PrintData(data)
}

func PrepareData(filename string) [][]string {
	strlist := make([]string, 0)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		strlist = append(strlist, scanner.Text())
	}

	var matrix [][]string
	for _, str := range strlist {
		matrix = append(matrix, strings.Fields(str))
	}

	return matrix
}

func DeleteDuplicates(data [][]string) [][]string {
	set := make(map[string]struct{})
	for i := 0; i < len(data); i++ {
		strItem := strings.ToLower(strings.Join(data[i], " "))
		if _, ok := set[strItem]; !ok {
			set[strItem] = struct{}{}
		}
	}
	// distinct
	dst := make([][]string, len(set))
	i := 0
	for key := range set {
		dst[i] = strings.Split(key, " ")
		i++
	}
	return dst
}

func Sort(data [][]string, k int, n, r bool) [][]string {

	if r && n {
		sort.Slice(data, func(i, j int) bool { return StrToInt(data[i][k-1]) > StrToInt(data[j][k-1]) })
	} else if r && !n {
		sort.Slice(data, func(i, j int) bool { return data[i][k-1] > data[j][k-1] })
	} else if !r && n {
		sort.Slice(data, func(i, j int) bool { return StrToInt(data[i][k-1]) < StrToInt(data[j][k-1]) })
	} else {
		sort.Slice(data, func(i, j int) bool { return data[i][k-1] < data[j][k-1] })
	}

	return data
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func PrintData(data [][]string) {
	for _, s1 := range data {
		for _, s2 := range s1 {
			fmt.Print(s2, " ")
		}
		fmt.Println()
	}
}
