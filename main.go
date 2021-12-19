package main

import (
	"fmt"
	"math/rand"
	"time"

	"algo/sort"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	fmt.Printf("%d\n", sort.Counting(arr))
}
