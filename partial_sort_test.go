package main

import (
	"math"
	"sort"
	"testing"

	csort "github.com/palletone/go-palletone/core/sort"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestPartialSortIntSlice(t *testing.T) {
	data := ints

	a := sort.IntSlice(data[0:])
	n := a.Len()

	m := 4
	p := sort.IntSlice(data[:m])

	csort.PartialSort(a, m)
	if !sort.IsSorted(p) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
		return
	}

	for i := m; i < n; i++ {
		if a.Less(i, m-1) {
			t.Errorf("sorted %v", ints)
			t.Errorf("   got %v", data)
		}
		return
	}
}

func TestPartialSortFloat64Slice(t *testing.T) {
	data := float64s

	a := sort.Float64Slice(data[0:])
	n := a.Len()

	m := 5
	p := sort.Float64Slice(data[:m])

	csort.PartialSort(a, m)
	if !sort.IsSorted(p) {
		t.Errorf("sorted %v", float64s)
		t.Errorf("   got %v", data)
		return
	}

	for i := m; i < n; i++ {
		if a.Less(i, m-1) {
			t.Errorf("sorted %v", ints)
			t.Errorf("   got %v", data)
			return
		}
	}
}

func TestPartialSortStringSlice(t *testing.T) {
	data := strings

	a := sort.StringSlice(data[0:])
	n := a.Len()

	m := 3
	p := sort.StringSlice(data[:m])

	csort.PartialSort(a, m)
	if !sort.IsSorted(p) {
		t.Errorf("sorted %v", strings)
		t.Errorf("   got %v", data)
		return
	}

	for i := m; i < n; i++ {
		if a.Less(i, m-1) {
			t.Errorf("sorted %v", ints)
			t.Errorf("   got %v", data)
			return
		}
	}
}
