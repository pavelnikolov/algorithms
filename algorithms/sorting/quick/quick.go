package quick

// Sort sorts the provided array using Quicksort algorithm
func Sort(arr []int) {
	var sort func(left int, right int)
	partition := func(left int, right int, pivot int) int {
		v := arr[pivot]
		right--
		arr[pivot], arr[right] = arr[right], arr[pivot]

		for i := left; i < right; i++ {
			if arr[i] <= v {
				arr[i], arr[left] = arr[left], arr[i]
				left++
			}
		}

		arr[left], arr[right] = arr[right], arr[left]
		return left
	}

	sort = func(left int, right int) {
		if left < right {
			pivot := (right + left) / 2
			pivot = partition(left, right, pivot)
			sort(left, pivot)
			sort(pivot+1, right)
		}
	}

	sort(0, len(arr))
}
