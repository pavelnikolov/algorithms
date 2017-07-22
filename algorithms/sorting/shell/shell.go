package shell

// Sort sorts an array using Shell sorting algorithm
func Sort(a []int) {
	increment := len(a) / 2
	for increment > 0 {
		for i := increment; i < len(a); i++ {
			j := i
			temp := a[i]

			for j >= increment && a[j-increment] > temp {
				a[j] = a[j-increment]
				j = j - increment
			}
			a[j] = temp
		}
		if increment == 2 {
			increment = 1
		} else {
			increment = int(increment * 5 / 11)
		}
	}
}

// Sort2 sorts the input array using shell sorting algorithm with predefined increments (gaps) list
func Sort2(a []int) {
	// http://www.cs.princeton.edu/~rs/shell/paperF.pdf
	// http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.14.9679&rep=rep1&type=pdf
	// sort gaps - the runtime performance depends on the gaps chosen and the input array; this is an unstable algorithm
	gaps := []int{262913, 65921, 16577, 4193, 1073, 281, 77, 23, 8, 1} // Theorem (Sedgewick, 1982)
	var x int
	for _, gap := range gaps {
		for i := gap; i < len(a); i++ {
			// add a[i] to the elements that have been gap sorted
			// save a[i] in temp and make a hole at position i
			x = a[i]
			// shift earlier gap-sorted elements up until the correct location for a[i] is found
			var j int
			for j = i; j >= gap && a[j-gap] > x; j -= gap {
				a[j] = a[j-gap]
			}
			// put temp (the original a[i]) in its correct location
			a[j] = x
		}
	}
}
