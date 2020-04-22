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

func sortMap(v interface{}) (r []keyValue, err error) {
	switch v.(type) {
	case map[string]int:
		for k, v := range v.(map[string]int) {
			r = append(r, keyValue{k, v})
		}
	case map[int]int:
		for k, v := range v.(map[int]int) {
			r = append(r, keyValue{k, v})
		}
	default:
		return nil, errors.New("Unsupported type")
	}
	sort.Slice(r, func(i, j int) bool {
		return r[i].Value > r[j].Value
	})
	return r, nil
}

func main() {
	sorted, err := sortMap(values)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i, v := range sorted {
		fmt.Printf("%s, %d\n", v.Key.(string), v.Value)
		if i >= 5 {
			fmt.Println("Top 5 vypsano")
			break
		}
	}

}
