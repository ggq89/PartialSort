package main

import (
	"math/rand"
	"time"
	"sort"
	"fmt"
)

const (
	//BIG_ARR_SIZE = 10*10*10*10*10*10*10*10
	BIG_ARR_SIZE = 10*4
	//RES_ARR_SIZE = 10*10*10
	RES_ARR_SIZE = 10
	//SplitPoint = 400
	SplitPoint =0
)

func fillArr(BigArr []int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < BIG_ARR_SIZE; i++ {
		BigArr[i] = rand.Intn(BIG_ARR_SIZE*10)
	}
}

type PartialSortFn func(data sort.Interface)

func testFn(fn PartialSortFn, BigArr [BIG_ARR_SIZE]int)  {
	fn(sort.Reverse(sort.IntSlice(BigArr[:])))
	fmt.Println(BigArr[:RES_ARR_SIZE])
}

func main() {
	var BigArr [BIG_ARR_SIZE]int //:= make([]int, BIG_ARR_SIZE, BIG_ARR_SIZE)
	fillArr(BigArr[:])
	fmt.Println(BigArr)

	testFn(solution1, BigArr)
	testFn(solution2, BigArr)
	testFn(solution3, BigArr)
	testFn(solution4, BigArr)
	testFn(solution5, BigArr)
	testFn(solution6, BigArr)
	testFn(solution7, BigArr)
}
