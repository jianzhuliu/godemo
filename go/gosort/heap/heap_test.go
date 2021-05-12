package heap

import (
	"gosort"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	maxCnt := 10000
	for i := 0; i < maxCnt; i++ {
		rand.Seed(time.Now().UnixNano())
		intSlice := gosort.GetRandIntSlice(maxCnt)
		heap := New(intSlice...)

		got := heap.Sort()
		sort.Ints(intSlice)
		if !reflect.DeepEqual(got, intSlice) {
			t.Fatalf("export %v,\nbut got %v\n", intSlice, got)
		}
	}
}

func TestSimple(t *testing.T) {
	intSlice := gosort.GetRandIntSlice(20)
	heap := New(intSlice...)
	t.Log("old:", intSlice)
	t.Log("size:", heap.Size())
	t.Log("data:", heap.Data())

	got := heap.Sort()
	t.Log("sort:", got)
	sort.Ints(intSlice)
	if !reflect.DeepEqual(got, intSlice) {
		t.Fatalf("export %v,\nbut got %v\n", intSlice, got)
	}
}
