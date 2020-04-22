package main

import (
	"fmt"
	"sort"
)

const topValuesCount = 3

var (
	values       map[string]int
	topValues    map[string]int // Seznam aktualne nejvyzsich hodnot
	topValuesMin int            // Aktualne nejmensi hodnota v "topValues"
)

type keyValue struct {
	Key   interface{}
	Value int
}

func init() {
	values = map[string]int{
		"A": 3,
		"B": 17,
		"C": 1,
		"D": 33,
		"E": 4,
		"F": 1,
		"G": 71,
		"H": 0,
		"I": 3,
	}
}

func sortMap(v interface{}) []keyValue {
	var s []keyValue
	for k, v := range v.(map[string]int) {
		s = append(s, keyValue{k, v})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Value > s[j].Value
	})
	return s
}

func main() {
	sortedX := sortMap(values)
	for _, v := range sortedX {
		fmt.Printf("%s, %d\n", v.Key.(string), v.Value)
	}

}
