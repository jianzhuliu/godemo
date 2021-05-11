package heap

import (
	"sort"
	"testing"
)

func TestHeap(t *testing.T) {
	intSlice := []int{8, 2, 6, 3, 1}
	heap := New(intSlice...)
	t.Log("old:", intSlice)
	t.Log("size:", heap.Size())
	t.Log("data:", heap.Data())
	size := heap.Size()
	for i := 0; i < size; i++ {
		t.Log("pop:", heap.Pop())
	}

	sort.Ints(intSlice)
	t.Log("sorted:", intSlice)
}
