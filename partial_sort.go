package main

import (
	"time"
	"fmt"
	"sort"
)

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
				symMerge(sort.Reverse(sort.IntSlice(ResArr)),0, SplitPoint, RES_ARR_SIZE)
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
	makeHeap(sort.Reverse(sort.IntSlice(ResArr)), 0, RES_ARR_SIZE)
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

	heapSort(sort.Reverse(sort.IntSlice(ResArr)), 0, RES_ARR_SIZE)
	//sort.Sort(sort.Reverse(sort.IntSlice(ResArr)))
}
