package main

import (
	"math/rand"
	"time"
	"sort"
	"fmt"
)

const (
	bigArrSize = 10*10*10*10*10*10
	resArrSize = 10*10*10
)

func fillSlice(data []int) {
	len := len(data)

	rand.Seed(time.Now().Unix())
	for i := 0; i < len; i++ {
		data[i] = rand.Int()
	}
}

type PartialSortFn func(data sort.Interface, m int)

func testFn(fn PartialSortFn, data [bigArrSize]int)  {
	start :=time.Now()
	fn(sort.Reverse(sort.IntSlice(data[:])), resArrSize)

	fmt.Println(time.Since(start))
}

func main() {
	var bigArr [bigArrSize]int
	fillSlice(bigArr[:])

	//testFn(solution1, bigArr)
	testFn(solution2, bigArr)
	testFn(solution3, bigArr)
	testFn(solution4, bigArr)
	testFn(solution5, bigArr)
	testFn(solution6, bigArr)
	testFn(solution7, bigArr)
}
