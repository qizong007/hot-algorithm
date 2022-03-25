package mysort

func UnguardedInsertSort(arr []int, l, r int) {
	idx := l

	// 找出最小值的index -> idx
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[idx] {
			idx = i
		}
	}

	// 把 idx 放在第一位（保证排序稳定性，所以冒泡交换）
	for idx > l {
		arr[idx], arr[idx-1] = arr[idx-1], arr[idx]
		idx--
	}

	// 下面就不需要每次都比较j是否大于l了（边界检测）
	// l+2开始，是因为[0]已经最小了
	for i := l + 2; i <= r; i++ {
		j := i
		for arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}
