package gosort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

//go test -v -run=TestHeap$
func TestHeap(t *testing.T) {
	maxCnt := 10000
	for i := 0; i < maxCnt; i++ {
		rand.Seed(time.Now().UnixNano())
		intSlice, copyData := GetRandIntSlice(maxCnt)

		Heap(sort.IntSlice(intSlice))
		sort.Ints(copyData)
		if !reflect.DeepEqual(intSlice, copyData) {
			t.Fatalf("export %v,\nbut got %v\n", copyData, intSlice)
		}
	}
}

//go test -v -run=TestHeapShow
func TestHeapShow(t *testing.T) {
	intSlice, copyData := GetRandIntSlice(20)

	t.Log("data:", intSlice)

	Heap(sort.IntSlice(intSlice))
	sort.Ints(copyData)

	if !reflect.DeepEqual(intSlice, copyData) {
		t.Fatalf("export %v,\nbut got %v\n", copyData, intSlice)
	}
}
