package main

import "sort"

// 对data中的数据进行heap排序
func HeapSort(data sort.Interface) {
	n := data.Len()
	makeHeap(data, 0, n)
	heapSort(data, 0, n)
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
// Reference from Go SDK 1.10.2 package sort.
// 对data[lo, hi)内的元素进行siftDown处理，以满足heap序列的要求
func siftDown(data sort.Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1 // 算出左子节点
		if child >= hi { // 判断是否超出范围
			break
		}
		// 找出值较大的子节点
		if child+1 < hi && data.Less(first+child, first+child+1) {
			child++
		}
		// 判断父节点是否小于最大的子节点，如果不小于则终止循环
		if !data.Less(first+root, first+child) {
			return
		}
		// 交换父节点和最大的子节点的值，继续对下一层做以上处理
		data.Swap(first+root, first+child)
		root = child
	}
}

// Build heap with greatest element at top.
// 将data[a, b)的数据排列为一个max-heap
func makeHeap(data sort.Interface, a, b int) {
	first := a
	hi := b - a

	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}
}

// 将一个max-heap进行堆排序
func heapSort(data sort.Interface, a, b int) {
	first := a
	hi := b - a

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		siftDown(data, first, i, first)
	}
}
