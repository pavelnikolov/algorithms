package selection

// Sort sorts the input array using the selection sort algorithm
func Sort(a []int) {
	n := len(a)
	var min int // the index of the minimal element
	for i := 0; i < n-1; i++ {
		min = i // assume that the first element in the array is the minimal
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}
}
