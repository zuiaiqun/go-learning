package test

import "utils"

func partion(data []int, start int, end int) int {
	key := data[start]
	i := start
	j := end
	for i < j {
		for i < j && data[j] >= key {
			j = j - 1
		}
		if i < j {
			data[i] = data[j]
			i += 1
		}
		for i < j && data[i] <= key {
			i = i + 1
		}
		if i < j {
			data[j] = data[i]
			j -= 1
		}
	}
	data[i] = key
	return i
}

func quick_sort(data []int, start int, end int) {
	if start >= end {
		return
	}
	pos := partion(data, start, end)
	quick_sort(data, 0, pos-1)
	quick_sort(data, pos+1, end)
}

func TestQuickSort() {
	var test_data = []int{1, 4, 6, 2, 3, 5, 9}
	quick_sort(test_data, 0, len(test_data)-1)
	utils.PrintArray(test_data)
	test_data = []int{1, 14, 6, 2, 8, 3, 5}
	quick_sort(test_data, 0, len(test_data)-1)
	utils.PrintArray(test_data)
}
