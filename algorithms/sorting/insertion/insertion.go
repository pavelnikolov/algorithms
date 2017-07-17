package insertion

// Sort sorts the input array using the insertion sort algorithm
func Sort(a []int) {
	var x, j int // tmp variable
	for i := 1; i < len(a); i++ {
		x = a[i]
		j = i - 1
		for ; j >= 0 && a[j] > x; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = x
	}
}
