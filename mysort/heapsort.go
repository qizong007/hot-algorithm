package mysort

func heapify(arr []int, n int, pos int) {
	if pos >= n {
		return
	}
	left := (pos << 1) + 1
	right := (pos << 1) + 2
	max := pos
	if left < n && arr[left] > arr[max] {
		max = left
	}
	if right < n && arr[right] > arr[max] {
		max = right
	}
	if max != pos {
		arr[max], arr[pos] = arr[pos], arr[max]
		heapify(arr, n, max)
	}
}

func buildHeap(arr []int) {
	last := len(arr) - 1
	parent := (last - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}
}

func HeapSort(arr []int) {
	buildHeap(arr)
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}
}
