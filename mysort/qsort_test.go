package mysort

import (
	"fmt"
	"hot-algorithm/util"
	"sort"
	"testing"
)

const ArrLen = 10000000

var arr []int

func init() {
	arr = util.GenerateIntList(ArrLen)
}

func quickSortCommon(quickSort func(arr []int, left int, right int)) {
	list := make([]int, ArrLen)
	copy(list, arr)
	util.ExecuteWithTime(func() {
		quickSort(list, 0, len(list)-1)
	})
	if !util.CheckIntSorted(list) {
		fmt.Println("Sort Fail!")
	}
}

func TestGoSort(t *testing.T) {
	list := make([]int, ArrLen)
	copy(list, arr)
	util.ExecuteWithTime(func() {
		sort.Ints(list)
	})
	if !util.CheckIntSorted(list) {
		fmt.Println("Sort Fail!")
	}
}

func TestHeapSort(t *testing.T) {
	list := make([]int, ArrLen)
	copy(list, arr)
	util.ExecuteWithTime(func() {
		HeapSort(list)
	})
	if !util.CheckIntSorted(list) {
		fmt.Println("Sort Fail!")
	}
}

func TestQuickSortV1(t *testing.T) {
	quickSortCommon(QuickSortV1)
}

func TestQuickSortV2(t *testing.T) {
	quickSortCommon(QuickSortV2)
}

func TestQuickSortV3(t *testing.T) {
	quickSortCommon(QuickSortV3)
}

func TestQuickSortV4(t *testing.T) {
	quickSortCommon(QuickSortV4)
}
