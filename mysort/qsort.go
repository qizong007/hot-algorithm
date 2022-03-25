package mysort

func QuickSortV1(arr []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := arr[left]
	i := left
	j := right
	for i < j {
		for i < j && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = pivot
	QuickSortV1(arr, left, i-1)
	QuickSortV1(arr, i+1, right)
}

// 获取三中值
func median(arr []int, l, r int) int {
	a := arr[l]
	b := arr[r]
	c := arr[(l+r)>>1]
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	return b
}

// QuickSortV2 V1 + 三中值 pivot
func QuickSortV2(arr []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := median(arr, left, right)
	i := left
	j := right
	for i < j {
		for i < j && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = pivot
	QuickSortV2(arr, left, i-1)
	QuickSortV2(arr, i+1, right)
}

// QuickSortV3 V2 + 单边递归、双指针、无监督
func QuickSortV3(arr []int, left int, right int) {
	for left < right {
		i := left
		j := right
		pivot := median(arr, left, right)
		for {
			for arr[i] < pivot {
				i++
			}
			for arr[j] > pivot {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
			if i > j {
				break
			}
		}
		QuickSortV3(arr, left, j)
		left = i
	}
}

// QuickSortV4 V3 + 插入排序
func QuickSortV4(arr []int, left int, right int) {
	for right-left > 16 {
		i := left
		j := right
		pivot := median(arr, left, right)
		for {
			for arr[i] < pivot {
				i++
			}
			for arr[j] > pivot {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
			if i > j {
				break
			}
		}
		QuickSortV4(arr, left, j)
		left = i
	}
	UnguardedInsertSort(arr, left, right)
}
