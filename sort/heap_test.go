package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestHeapBuild(t *testing.T) {
	arr := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	HeapSort(arr)
	fmt.Println(arr, sort.IntsAreSorted(arr))
}
