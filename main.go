package main

import (
	"errors"
	"fmt"
	"sort"
)

const topValuesCount = 3

var (
	values       map[string]int
	topValues    map[string]int // Seznam aktualne nejvyzsich hodnot
	topValuesMin int            // Aktualne nejmensi hodnota v "topValues"
)

type keyStringValue struct {
	Key   string
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

func sortMap(v interface{}) (interface{}, error) {
	switch v.(type) {
	case map[string]int:
		var s []keyStringValue
		for k, v := range v.(map[string]int) {
			s = append(s, keyStringValue{k, v})
		}
		sort.Slice(s, func(i, j int) bool {
			return s[i].Value > s[j].Value
		})
		return s, nil
	default:
		return nil, errors.New("Unsupported type")
	}
}

func main() {
	sortedX, _ := sortMap(values)
	for _, v := range sortedX.([]keyStringValue) {
		fmt.Printf("%s, %d\n", v.Key, v.Value)
	}

}
