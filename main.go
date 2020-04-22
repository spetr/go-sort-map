package main

import (
	"errors"
	"fmt"
	"sort"
)

const topValuesCount = 3

var (
	valuesString map[string]int
	valuesInt    map[int]int
)

type keyValue struct {
	Key   interface{}
	Value int
}

func init() {
	valuesString = map[string]int{
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
	valuesInt = map[int]int{
		1: 53,
		2: 9,
		3: 66,
		4: 0,
		5: 11,
		6: 27,
		7: 9,
		8: 44,
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
	fmt.Println("Key is String:")
	sortedString, err := sortMap(valuesString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range sortedString {
		fmt.Printf("%s, %d\n", v.Key.(string), v.Value)
	}

	fmt.Println("Key is Int:")
	sortedInt, err := sortMap(valuesInt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range sortedInt {
		fmt.Printf("%d, %d\n", v.Key.(int), v.Value)
	}

}
