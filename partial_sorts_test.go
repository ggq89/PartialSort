package main

import (
	"math"
	"sort"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(),
	math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func IsPartialSorted(data sort.Interface, m int) bool {
	mid := m-1
	for i := mid; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}

	n := data.Len()
	for i := m; i < n; i++ {
		if data.Less(i, mid) {
			return false
		}
	}

	return true
}

func testPartialSortIntSlice(t *testing.T, fn PartialSortFn) {
	data := ints
	a := sort.IntSlice(data[0:])

	m := 4
	fn(a, m)

	if !IsPartialSorted(a, m) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func testPartialSortFloat64Slice(t *testing.T, fn PartialSortFn) {
	data := float64s
	a := sort.Float64Slice(data[0:])

	m := 5
	fn(a, m)

	if !IsPartialSorted(a, m) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func testPartialSortStringSlice(t *testing.T, fn PartialSortFn) {
	data := strings
	a := sort.StringSlice(data[0:])

	m := 3
	fn(a, m)

	if !IsPartialSorted(a, m) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func testPartialSort(t *testing.T, fn PartialSortFn) {
	testPartialSortIntSlice(t, fn)
	testPartialSortFloat64Slice(t, fn)
	testPartialSortStringSlice(t, fn)
}

func TestPartialSorts(t *testing.T) {
	testPartialSort(t, solution1)
	testPartialSort(t, solution2)
	testPartialSort(t, solution3)
	testPartialSort(t, solution4)
	testPartialSort(t, solution5)
	testPartialSort(t, solution6)
	testPartialSort(t, solution7)
}
