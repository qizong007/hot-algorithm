package mysort

import "fmt"

func merge(arr []int, left int, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	merge(arr, left, mid)
	merge(arr, mid+1, right)
	leftSize := mid - left + 1
	rightSize := right - mid
	lArr := make([]int, leftSize)
	rArr := make([]int, rightSize)
	for i := left; i <= mid; i++ {
		lArr[i-left] = arr[i]
	}
	for i := mid + 1; i <= right; i++ {
		rArr[i-mid-1] = arr[i]
	}
	i := 0
	j := 0
	cnt := left
	for i < leftSize && j < rightSize {
		if lArr[i] < rArr[j] {
			arr[cnt] = lArr[i]
			cnt++
			i++
		} else {
			arr[cnt] = rArr[j]
			cnt++
			j++
		}
	}
	for i < leftSize {
		arr[cnt] = lArr[i]
		cnt++
		i++
	}
	for j < rightSize {
		arr[cnt] = rArr[j]
		cnt++
		j++
	}
}

func mergesort(arr []int) {
	n := len(arr)
	merge(arr, 0, n-1)
}

func main() {
	arr := []int{123, 4312, 132, -31, 0, 1}
	fmt.Println(arr)
	mergesort(arr)
	fmt.Println(arr)
}
