package bubble

// Sort sorts the input array using the bubble sort algorithm
func Sort(a []int) {
	var swapped bool
	for i := len(a) - 1; i > 0; i-- {
		swapped = false
		for j := 1; j < i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
