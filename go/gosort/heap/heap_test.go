package heap

import (
	"gosort"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

//go test -v -run=TestSort
func TestSort(t *testing.T) {
	maxCnt := 10000
	for i := 0; i < maxCnt; i++ {
		rand.Seed(time.Now().UnixNano())
		intSlice, copyData := gosort.GetRandIntSlice(maxCnt)
		heap := New(intSlice...)

		got := heap.Sort()
		sort.Ints(copyData)
		if !reflect.DeepEqual(got, copyData) {
			t.Fatalf("export %v,\nbut got %v\n", copyData, got)
		}
	}
}

//go test -v -run=TestSimple
func TestSimple(t *testing.T) {
	intSlice, copyData := gosort.GetRandIntSlice(20)
	heap := New(intSlice...)
	t.Log("old:", intSlice)
	t.Log("size:", heap.Size())
	t.Log("data:", heap.Data())

	got := heap.Sort()
	t.Log("sort:", got)
	sort.Ints(copyData)
	if !reflect.DeepEqual(got, copyData) {
		t.Fatalf("export %v,\nbut got %v\n", copyData, got)
	}
}
