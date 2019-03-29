package main

import (
	"time"
	"fmt"
	"sort"
)

func solution1(data sort.Interface) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	for i := 0; i < RES_ARR_SIZE; i++ {
		idx := i
		// 选出[i, data.Len)中最小的元素
		for j := i+1; j < BIG_ARR_SIZE; j++ {
			if data.Less(j, idx) {
				idx = j
			}
		}
		data.Swap(i, idx)
	}
}

func solution2(data sort.Interface) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	sort.Sort(data)
}

func solution3(data sort.Interface) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//标记是否发生过交换
	bExchanged := true
	var idx int

	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		//如果上一轮发生过交换
		if bExchanged {
			//找出ResArr中最大的元素
			idx = 0
			for j := idx + 1; j < RES_ARR_SIZE; j++ {
				if data.Less(idx, j) {
					idx = j
				}
			}
		}

		bExchanged = false
		//这个后续元素比ResArr中最大的元素小，则替换
		if data.Less(i, idx) {
			bExchanged = true
			data.Swap(i, idx)
		}
	}

	intervalSort(data, 0, RES_ARR_SIZE)
}

func solution4(data sort.Interface)  {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//排序
	intervalSort(data, 0, RES_ARR_SIZE)

	//最大元素索引
	maxElemIdx := RES_ARR_SIZE - 1
	//可能产生交换的区域的最小索引
	zoneBeginIdx := maxElemIdx

	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		//这个后续元素比ResArr中最大的元素小，则替换
		if data.Less(i, maxElemIdx) {
			data.Swap(maxElemIdx, i)
			if zoneBeginIdx == maxElemIdx && zoneBeginIdx > 0 {
				zoneBeginIdx--
			}

			//查找最大元素
			idx := zoneBeginIdx
			for j := idx + 1; j < RES_ARR_SIZE; j++ {
				if data.Less(idx, j) {
					idx = j
				}
			}
			maxElemIdx = idx
		}
	}

	intervalSort(data, 0, RES_ARR_SIZE)
}

func solution5(data sort.Interface)  {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//排序
	intervalSort(data, 0, RES_ARR_SIZE)

	//最大元素索引
	maxElemIdx := RES_ARR_SIZE - 1
	//可能产生交换的区域的最小索引
	zoneBeginIdx := maxElemIdx

	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		//这个后续元素比ResArr中最大的元素小，则替换
		if data.Less(i, maxElemIdx) {
			data.Swap(maxElemIdx, i)
			if zoneBeginIdx == maxElemIdx && zoneBeginIdx > 0 {
				zoneBeginIdx--
			}

			//太多杂乱元素的时候排序
			if zoneBeginIdx < SplitPoint {
				intervalSort(data, 0, RES_ARR_SIZE)
				maxElemIdx = RES_ARR_SIZE - 1
				zoneBeginIdx = maxElemIdx
				continue
			}

			//查找最大元素
			idx := zoneBeginIdx
			for j := idx + 1; j < RES_ARR_SIZE; j++ {
				if data.Less(idx, j) {
					idx = j
				}
			}
			maxElemIdx = idx
		}
	}

	intervalSort(data, 0, RES_ARR_SIZE)
}

func solution6(data sort.Interface)  {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//排序
	intervalSort(data, 0, RES_ARR_SIZE)

	//最大元素索引
	maxElemIdx := RES_ARR_SIZE - 1
	//可能产生交换的区域的最小索引
	zoneBeginIdx := maxElemIdx

	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		//这个后续元素比ResArr中最大的元素小，则替换
		if data.Less(i, maxElemIdx) {
			data.Swap(maxElemIdx, i)
			if zoneBeginIdx == maxElemIdx && zoneBeginIdx > 0 {
				zoneBeginIdx--
			}

			//太多杂乱元素的时候排序
			if zoneBeginIdx < SplitPoint {
				intervalSort(data, 0, RES_ARR_SIZE)
				symMerge(data,0, SplitPoint, RES_ARR_SIZE)
				maxElemIdx = RES_ARR_SIZE - 1
				zoneBeginIdx = maxElemIdx
				continue
			}

			//查找最大元素
			idx := zoneBeginIdx
			for j := idx + 1; j < RES_ARR_SIZE; j++ {
				if data.Less(idx, j) {
					idx = j
				}
			}
			maxElemIdx = idx
		}
	}

	intervalSort(data, 0, RES_ARR_SIZE)
}

func solution7(data sort.Interface) {
	defer func(start time.Time) {
		fmt.Println(time.Since(start))
	}(time.Now())

	//建max-heap
	makeHeap(data, 0, RES_ARR_SIZE)
	minElemIdx := 0

	//遍历后续的元素
	for i := RES_ARR_SIZE; i < BIG_ARR_SIZE; i++ {
		if data.Less(i, minElemIdx) {
			//当这个后续元素比ResArr中最大的元素小，则替换
			data.Swap(i, minElemIdx)
			// 重新调整max-heap
			siftDown(data, minElemIdx, RES_ARR_SIZE, minElemIdx)
		}
	}

	// 对max-heap进行对排序
	heapsort(data, 0, RES_ARR_SIZE)
}
