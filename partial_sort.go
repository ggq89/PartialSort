package main

import (
	"sort"
)

func solution1(data sort.Interface, m int) {
	len := data.Len()
	
	for i := 0; i < m; i++ {
		idx := i
		// 选出[i, data.Len)中最小的元素
		for j := i+1; j < len; j++ {
			if data.Less(j, idx) {
				idx = j
			}
		}
		data.Swap(i, idx)
	}
}

func solution2(data sort.Interface, m int) {
	sort.Sort(data)
}

func solution3(data sort.Interface, m int) {
	//标记是否发生过交换
	bExchanged := true
	var idx int	// 后续所有元素中最大元素的索引

	//遍历后续的元素
	len := data.Len()
	for i := m; i < len; i++ {
		//如果上一轮发生过交换
		if bExchanged {
			//找出ResArr中最大的元素
			idx = 0
			for j := idx + 1; j < m; j++ {
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

	intervalSort(data, 0, m)
}

func solution4(data sort.Interface, m int)  {
	//排序
	intervalSort(data, 0, m)

	//最大元素索引
	maxElemIdx := m - 1
	//可能产生交换的区域的最小索引
	zoneBeginIdx := maxElemIdx

	//遍历后续的元素
	len := data.Len()
	for i := m; i < len; i++ {
		//这个后续元素比ResArr中最大的元素小，则替换
		if data.Less(i, maxElemIdx) {
			data.Swap(maxElemIdx, i)
			if zoneBeginIdx == maxElemIdx && zoneBeginIdx > 0 {
				zoneBeginIdx--
			}

			//查找最大元素
			idx := zoneBeginIdx
			for j := idx + 1; j < m; j++ {
				if data.Less(idx, j) {
					idx = j
				}
			}
			maxElemIdx = idx
		}
	}

	intervalSort(data, 0, m)
}

func solution5(data sort.Interface, m int)  {
	//排序
	intervalSort(data, 0, m)

	//最大元素索引
	maxElemIdx := m - 1
	//可能产生交换的区域的最小索引
	zoneBeginIdx := maxElemIdx

	//遍历后续的元素
	len := data.Len()
	for i := m; i < len; i++ {
		//这个后续元素比ResArr中最大的元素小，则替换
		if data.Less(i, maxElemIdx) {
			data.Swap(maxElemIdx, i)
			if zoneBeginIdx == maxElemIdx && zoneBeginIdx > 0 {
				zoneBeginIdx--
			}

			//太多杂乱元素的时候排序
			if zoneBeginIdx < SplitPoint {
				intervalSort(data, 0, m)
				maxElemIdx = m - 1
				zoneBeginIdx = maxElemIdx
				continue
			}

			//查找最大元素
			idx := zoneBeginIdx
			for j := idx + 1; j < m; j++ {
				if data.Less(idx, j) {
					idx = j
				}
			}
			maxElemIdx = idx
		}
	}

	intervalSort(data, 0, m)
}

const (
	SplitPoint = 400
)

func solution6(data sort.Interface, m int)  {
	//排序
	intervalSort(data, 0, m)

	//最大元素索引
	maxElemIdx := m - 1
	//可能产生交换的区域的最小索引
	zoneBeginIdx := maxElemIdx

	//遍历后续的元素
	len := data.Len()
	for i := m; i < len; i++ {
		//这个后续元素比ResArr中最大的元素小，则替换
		if data.Less(i, maxElemIdx) {
			data.Swap(maxElemIdx, i)
			if zoneBeginIdx == maxElemIdx && zoneBeginIdx > 0 {
				zoneBeginIdx--
			}

			//太多杂乱元素的时候排序
			if zoneBeginIdx < SplitPoint {
				intervalSort(data, 0, m)
				symMerge(data,0, zoneBeginIdx+1, m)
				maxElemIdx = m - 1
				zoneBeginIdx = maxElemIdx
				continue
			}

			//查找最大元素
			idx := zoneBeginIdx
			for j := idx + 1; j < m; j++ {
				if data.Less(idx, j) {
					idx = j
				}
			}
			maxElemIdx = idx
		}
	}

	intervalSort(data, 0, m)
}

func solution7(data sort.Interface, m int) {
	PartialSort(data, m)
}

// PartialSort, Rearranges elements such that the range [0, m)
// contains the sorted m smallest elements in the range [first, data.Len).
// The order of equal elements is not guaranteed to be preserved.
// The order of the remaining elements in the range [m, data.Len) is unspecified.
func PartialSort(data sort.Interface, m int) {
	//建max-heap
	makeHeap(data, 0, m)
	minElemIdx := 0

	//遍历后续的元素
	len := data.Len()
	for i := m; i < len; i++ {
		if data.Less(i, minElemIdx) {
			//当这个后续元素比ResArr中最大的元素小，则替换
			data.Swap(i, minElemIdx)
			// 重新调整max-heap
			siftDown(data, minElemIdx, m, minElemIdx)
		}
	}

	// 对max-heap进行对排序
	heapsort(data, 0, m)
}
