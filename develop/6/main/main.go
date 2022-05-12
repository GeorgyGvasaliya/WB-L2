package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f := flag.String("f", "1", "fields")
	d := flag.String("d", "\t", "delimiter")
	s := flag.Bool("s", false, "separated")
	filename := flag.String("fn", "file.txt", "filename")

	flag.Parse()

	//  read data and transform it to matrix
	data := PrepareData(*filename, *d)

	// columns
	clms := StrFflagToIntList(*f)

	if *s {
		data = HandeS(data)
	}

	PrintFlagf(data, clms)

}

func PrepareData(filename string, del string) [][]string {
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
		matrix = append(matrix, strings.Split(str, del))
	}

	return matrix
}

// handle flag f -> list of columns
func StrFflagToIntList(str string) []int {
	clm := make([]int, 0)
	lststring := strings.Split(str, ",")
	for _, s := range lststring {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		clm = append(clm, i)
	}
	return clm
}

func crmMatrix(data []string, sep string) [][]string {
	var matrix [][]string
	for _, str := range data {
		matrix = append(matrix, strings.Split(str, sep))
	}
	return matrix
}

func HandeS(matr [][]string) [][]string {

	newmatr := [][]string{}
	for _, str := range matr {
		//there is sep if more than 1 elem
		if len(str) > 1 {
			newmatr = append(newmatr, str)
		}
	}
	matr = newmatr

	return matr
}

func PrintFlagf(matr [][]string, lst []int) {
	for _, str := range matr {
		// no sep print
		if len(str) == 1 {
			fmt.Print(str[0] + ":")
			fmt.Println()
			continue
		}

		for _, v := range lst {
			fmt.Print(str[v-1] + ":")
		}
		fmt.Println()
	}
}
