package main

import (
	"fmt"
	"sort"
)

func main() {

	//map的排序
	map1 := make(map[int]int)
	map1[1] = 90
	map1[20] = 88
	map1[10] = 11
	map1[3] = 2
	fmt.Println(map1)

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println("pre-keys", keys)
	sort.Ints(keys)
	fmt.Println("keys", keys)

	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}
}