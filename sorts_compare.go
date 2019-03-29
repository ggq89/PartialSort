package main

import (
	"math/rand"
	"time"
)

const (
	BIG_ARR_SIZE = 10*10*10*10*10*10*10*10
	//BIG_ARR_SIZE = 10*4
	//RES_ARR_SIZE = 10*10*10
	RES_ARR_SIZE = 10*10
	//SplitPoint = 400
	SplitPoint =0
)

var (
	BigArr = make([]int, BIG_ARR_SIZE, BIG_ARR_SIZE)
	ResArr []int
)

func fillArr() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < BIG_ARR_SIZE; i++ {
		BigArr[i] = rand.Intn(BIG_ARR_SIZE*10)
	}
}

type PartialSortFn func(BigArr, ResArr []int)

func testFn(fn PartialSortFn, BigArr []int)  {
	ResArr = make([]int, RES_ARR_SIZE, RES_ARR_SIZE)
	fn(BigArr, ResArr)
	//fmt.Println(ResArr)
}

func main() {
	fillArr()
	//fmt.Println(BigArr)

	//testFn(solution1, BigArr)
	//testFn(solution2, BigArr)
	testFn(solution3, BigArr)
	testFn(solution4, BigArr)
	testFn(solution5, BigArr)
	testFn(solution6, BigArr)
	testFn(solution7, BigArr)
}
