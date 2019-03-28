package main

import (
	"math/rand"
	"time"
	"fmt"
	"sort"
)

const (
	BIG_ARR_SIZE = 10*10*10*10*10*10
	//BIG_ARR_SIZE = 10*4
	RES_ARR_SIZE = 10*10*10
	//RES_ARR_SIZE = 10
	SplitPoint = 400
	//SplitPoint =0
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

func solution1(BigArr, ResArr []int) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	for i := 0; i < RES_ARR_SIZE; i++ {
		idx := i
		for j := i+1; j < BIG_ARR_SIZE; j++ {
			if BigArr[j] > BigArr[idx] {
				idx = j
			}
		}
		ResArr[i] = BigArr[idx]
		BigArr[i], BigArr[idx] = BigArr[idx], BigArr[i]
	}
}

func copyArr(desArr, sourArr []int, len int) {
	for i := 0; i<len; i++ {
		desArr[i] = sourArr[i]
	}
}

func solution2(BigArr, ResArr []int) {
	cpBigArr :=make([]int, 0, BIG_ARR_SIZE)
	cpBigArr = append(cpBigArr, BigArr...)

	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	sort.Sort(sort.Reverse(sort.IntSlice(cpBigArr)))
	copyArr(ResArr, cpBigArr, RES_ARR_SIZE)
}

func solution3(BigArr, ResArr []int) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//取最前面的RES_ARR_SIZE个
	copyArr(ResArr, BigArr, RES_ARR_SIZE)

	//标记是否发生过交换
	bExchanged := true

	var idx int
	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		//如果上一轮发生过交换
		if bExchanged {
			//找出ResArr中最小的元素
			idx = 0
			for j := idx + 1; j < RES_ARR_SIZE; j++ {
				if ResArr[j] > ResArr[idx] {
					idx = j
				}
			}
		}

		bExchanged = false
		//这个后续元素比ResArr中最小的元素大，则替换。
		if BigArr[i] > ResArr[idx] {
			bExchanged = true
			ResArr[idx] = BigArr[i]
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
}

func solution4(BigArr, ResArr []int)  {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//取最前面的RES_ARR_SIZE个
	copyArr(ResArr, BigArr, RES_ARR_SIZE)
	//排序
	sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
	//最小元素索引
	MinElemIdx := RES_ARR_SIZE - 1
	//可能产生交换的区域的最小索引
	ZoneBeginIdx := MinElemIdx
	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		//这个后续元素比ResArr中最小的元素大，则替换。
		if BigArr[i] > ResArr[MinElemIdx] {
			ResArr[MinElemIdx] = BigArr[i]
			if ZoneBeginIdx == MinElemIdx && ZoneBeginIdx > 0 {
				ZoneBeginIdx--
			}
			//查找最小元素
			idx := ZoneBeginIdx
			for j := idx + 1; j < RES_ARR_SIZE; j++ {
				if ResArr[j] > ResArr[idx] {
					idx = j
				}
			}
			MinElemIdx = idx
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
}

func solution5(BigArr, ResArr []int)  {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//取最前面的RES_ARR_SIZE个
	copyArr(ResArr, BigArr, RES_ARR_SIZE)
	//排序
	sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
	//最小元素索引
	MinElemIdx := RES_ARR_SIZE - 1
	//可能产生交换的区域的最小索引
	ZoneBeginIdx := MinElemIdx
	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		//这个后续元素比ResArr中最小的元素大，则替换。
		if BigArr[i] > ResArr[MinElemIdx] {
			ResArr[MinElemIdx] = BigArr[i]
			if ZoneBeginIdx == MinElemIdx && ZoneBeginIdx > 0 {
				ZoneBeginIdx--
			}

			//太多杂乱元素的时候排序
			if ZoneBeginIdx < SplitPoint {
				sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
				MinElemIdx = RES_ARR_SIZE - 1
				ZoneBeginIdx = MinElemIdx
				continue
			}
			
			//查找最小元素
			idx := ZoneBeginIdx
			for j := idx + 1; j < RES_ARR_SIZE; j++ {
				if ResArr[j] > ResArr[idx] {
					idx = j
				}
			}
			MinElemIdx = idx
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
}

//merge sorted sub-array in[p~q] and in[q+1~r] in-place, 0 <= p <= q < r= < in.Len()-1
func merge(in sort.Interface, p, q, r int) {
	i := p
	j := q + 1
	for k := p; k < r; k++ {
		if j > r {
			break
		}

		if !in.Less(i, j) {
			for m := j; m > i; m-- {
				in.Swap(m, m-1)
			}
			j++
		}
		i++
	}
}

func solution6(BigArr, ResArr []int)  {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//取最前面的RES_ARR_SIZE个
	copyArr(ResArr, BigArr, RES_ARR_SIZE)
	//排序
	sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
	//最小元素索引
	MinElemIdx := RES_ARR_SIZE - 1
	//可能产生交换的区域的最小索引
	ZoneBeginIdx := MinElemIdx
	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		//这个后续元素比ResArr中最小的元素大，则替换。
		if BigArr[i] > ResArr[MinElemIdx] {
			ResArr[MinElemIdx] = BigArr[i]
			if ZoneBeginIdx == MinElemIdx && ZoneBeginIdx > 0 {
				ZoneBeginIdx--
			}

			//太多杂乱元素的时候排序
			if ZoneBeginIdx < SplitPoint {
				sort.Sort(sort.Reverse(sort.IntSlice(ResArr[SplitPoint:RES_ARR_SIZE])))
				merge(sort.Reverse(sort.IntSlice(ResArr)),0, ZoneBeginIdx, RES_ARR_SIZE-1)
				MinElemIdx = RES_ARR_SIZE - 1
				ZoneBeginIdx = MinElemIdx
				continue
			}

			//查找最小元素
			idx := ZoneBeginIdx
			for j := idx + 1; j < RES_ARR_SIZE; j++ {
				if ResArr[j] > ResArr[idx] {
					idx = j
				}
			}
			MinElemIdx = idx
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
}

func solution7(BigArr, ResArr []int) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//取最前面的RES_ARR_SIZE个
	copyArr(ResArr, BigArr, RES_ARR_SIZE)
	//建min-heap
	makeHeap(sort.Reverse(sort.IntSlice(ResArr)))
	minElemIdx := 0

	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		if BigArr[i] > ResArr[minElemIdx] {
			//当这个后续元素比ResArr中最小的元素大，则替换
			ResArr[minElemIdx] = BigArr[i]
			// 重新调整min-heap
			siftDown(sort.Reverse(sort.IntSlice(ResArr)), minElemIdx, RES_ARR_SIZE, minElemIdx)
		}
	}

	heapSort(sort.Reverse(sort.IntSlice(ResArr)))
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(data sort.Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		if !data.Less(first+root, first+child) {
			return
		}
		data.Swap(first+root, first+child)
		root = child
	}
}

// Build heap with greatest element at top.
func makeHeap(data sort.Interface) {
	first := 0
	hi := data.Len()

	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}
}

func heapSort(data sort.Interface) {
	first := 0
	lo := 0
	hi := data.Len()

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, lo, i, first)
	}
}

func solution8(BigArr, ResArr []int) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//取最前面的RES_ARR_SIZE个
	copyArr(ResArr, BigArr, RES_ARR_SIZE)
	//建min-heap
	sort.Sort(sort.IntSlice(ResArr))
	minElemIdx := 0

	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		if BigArr[i] > ResArr[minElemIdx] {
			//当这个后续元素比ResArr中最小的元素大，则替换
			ResArr[minElemIdx] = BigArr[i]
			// 重新调整min-heap
			siftDown(sort.Reverse(sort.IntSlice(ResArr)), minElemIdx, RES_ARR_SIZE, minElemIdx)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
}

func solution9(BigArr, ResArr []int) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//取最前面的RES_ARR_SIZE个
	copyArr(ResArr, BigArr, RES_ARR_SIZE)
	//建min-heap
	makeHeap(sort.Reverse(sort.IntSlice(ResArr)))
	minElemIdx := 0

	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		if BigArr[i] > ResArr[minElemIdx] {
			//当这个后续元素比ResArr中最小的元素大，则替换
			ResArr[minElemIdx] = BigArr[i]
			// 重新调整min-heap
			siftDown(sort.Reverse(sort.IntSlice(ResArr)), minElemIdx, RES_ARR_SIZE, minElemIdx)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
}

func solution10(BigArr, ResArr []int) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//取最前面的RES_ARR_SIZE个
	copyArr(ResArr, BigArr, RES_ARR_SIZE)
	//建min-heap
	sort.Sort(sort.IntSlice(ResArr))
	minElemIdx := 0

	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		if BigArr[i] > ResArr[minElemIdx] {
			//当这个后续元素比ResArr中最小的元素大，则替换
			ResArr[minElemIdx] = BigArr[i]
			// 重新调整min-heap
			siftDown(sort.Reverse(sort.IntSlice(ResArr)), minElemIdx, RES_ARR_SIZE, minElemIdx)
		}
	}

	heapSort(sort.Reverse(sort.IntSlice(ResArr)))
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

	testFn(solution1, BigArr)
	testFn(solution2, BigArr)
	testFn(solution3, BigArr)
	testFn(solution4, BigArr)
	testFn(solution5, BigArr)
	testFn(solution6, BigArr)
	testFn(solution7, BigArr)
	testFn(solution8, BigArr)
	testFn(solution9, BigArr)
	testFn(solution10, BigArr)
}
