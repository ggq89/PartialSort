package main

import (
	"math/rand"
	"time"
	"sort"
	"fmt"
)

const (
	BIG_ARR_SIZE = 10*10*10*10*10*10
	RES_ARR_SIZE = 10*10*10
	//BIG_ARR_SIZE = 10*4
	//RES_ARR_SIZE = 10
)

func fillSlice(data []int) {
	len := len(data)
	n := len*10

	rand.Seed(time.Now().Unix())
	for i := 0; i < len; i++ {
		data[i] = rand.Intn(n)
	}
}

type PartialSortFn func(data sort.Interface, m int)

func testFn(fn PartialSortFn, BigArr [BIG_ARR_SIZE]int, m int)  {
	start :=time.Now()
	fn(sort.Reverse(sort.IntSlice(BigArr[:])), m)

	fmt.Println(time.Since(start))
	//fmt.Println(BigArr[:m])
}

func main() {
	var BigArr [BIG_ARR_SIZE]int
	fillSlice(BigArr[:])
	//fmt.Println(BigArr)

	//testFn(solution1, BigArr, RES_ARR_SIZE)
	//testFn(solution2, BigArr, RES_ARR_SIZE)
	testFn(solution3, BigArr, RES_ARR_SIZE)
	testFn(solution4, BigArr, RES_ARR_SIZE)
	testFn(solution5, BigArr, RES_ARR_SIZE)
	testFn(solution6, BigArr, RES_ARR_SIZE)
	testFn(solution7, BigArr, RES_ARR_SIZE)
}
