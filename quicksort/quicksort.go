package quicksort

// QuickSort 快速排序算法
func QuickSort(values []int, left, right int) {

	temp := values[left]

	p := left

	i, j := left, right

	for i <= j {

		for j >= p && values[j] >= temp {
			j--
		}

		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp

	if p-left > 1 {
		QuickSort(values, left, p-1)
	}

	if right-p > 1 {
		QuickSort(values, p+1, right)
	}
}
