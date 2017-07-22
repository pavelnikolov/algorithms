package merge

// Sort sorts the input array using Merge Sort algorithm (top to bottom)
func Sort(a []int) {
	if len(a) < 2 {
		return
	}
	mid := len(a) / 2

	Sort(a[:mid])
	Sort(a[mid:])

	// if the greatest element in the left half is less than or equal to the smallest element in the right half
	if a[mid-1] <= a[mid] {
		return
	}

	copy(a[:], Merge(a[:mid], a[mid:]))
}

// Merge expects two sorted arrays as input, it merges them and returns a soreted array as a result
func Merge(a1, a2 []int) []int {
	var i, j int
	n := len(a1) + len(a2)
	res := make([]int, n)

	for i+j < n {
		if i == len(a1) {
			copy(res[i+j:], a2[j:])
			break
		}
		if j == len(a2) {
			copy(res[i+j:], a1[i:])
			break
		}
		if a1[i] < a2[j] {
			res[i+j] = a1[i]
			i++
		} else {
			res[i+j] = a2[j]
			j++
		}
	}
	return res
}
