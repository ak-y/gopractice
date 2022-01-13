package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"algo/search"
	"algo/sort"
)

var (
	sortBy   string
	searchBy string
	sorts    []string = []string{"comb", "counting", "insertion", "select", "shell"}
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.StringVar(&sortBy, "sortBy", "shell", "sort-algorithm you want to use")
	flag.StringVar(&searchBy, "searchBy", "linear", "search-algorithm you want to use")
	flag.Parse()
}

func main() {
	if flag.NArg() != 1 {
		os.Exit(1)
	}

	do := flag.Arg(0)
	if do == "sort" {
		dosort()
	} else if do == "search" {
		dosearch()
	} else {
		os.Exit(1)
	}
}

func dosort() {
	// print sort-algorithm name
	correctName := false
	for _, v := range sorts {
		if strings.ToLower(sortBy) == v {
			fmt.Printf("%s-sort\n", v)
			correctName = true
		}
	}
	if !(correctName) {
		fmt.Println("shell-sort")
	}

	// ソート元の配列を作る
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	fmt.Println("before =>", arr)

	var ans []int
	switch strings.ToLower(sortBy) {
	case "comb":
		ans = sort.Comb(arr)
	case "counting":
		ans = sort.Counting(arr)
	case "insertion":
		ans = sort.Insertion(arr)
	case "select":
		ans = sort.Select(arr)
	default:
		ans = sort.Shell(arr)
	}

	fmt.Println("after  =>", ans)
}

func dosearch() {
	arr := []int{0, 1, 5, 7, 9, 11, 15, 20, 24}
	target := 11

	var ans int
	switch strings.ToLower(searchBy) {
	case "linear":
		fmt.Println("linear-search")
		ans = search.Linear(arr, target)
	case "iterbinary":
		fmt.Println("iterbinary-search")
		ans = search.IterBinary(arr, target)
	default:
		fmt.Println("binary-search")
		ans = search.Binary(arr, target)
	}

	fmt.Println(arr)
	fmt.Printf("target=%d\n", target)
	fmt.Println("target's index =>", ans)
}
