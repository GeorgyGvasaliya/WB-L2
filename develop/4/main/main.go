package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	in := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}

	fmt.Println(in)
	fmt.Println(AnagramMap(in))
}

func AnagramMap(sl []string) map[string][]string {
	for i := range sl {
		sl[i] = strings.ToLower(sl[i])
	}
	distinctSl := deleteDuplicates(sl)
	tempM := make(map[string][]string)

	for _, v := range distinctSl {
		sorted := []rune(v)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})
		sortedS := string(sorted)

		tempM[sortedS] = append(tempM[sortedS], v)
	}

	resultM := make(map[string][]string)

	for _, v := range tempM {
		if len(v) > 1 {
			resultM[v[0]] = v
		}
	}

	return resultM

}

func deleteDuplicates(sl []string) []string {
	m := make(map[string]bool)
	result := make([]string, 0)

	for _, v := range sl {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}
