package array

// IsRotation checks if b contains a rotation of array a
// a := []int{1, 2, 3, 4, 5, 6, 7}
// b := []int{4, 5, 6, 7, 1, 2, 3},
//            0  1  2  3  4  5  6
func IsRotation(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	// approach 1
	// var tables = make(map[int]int)
	// for i := 0; i < len(a); i++ {
	// 	tables[a[i]] = i
	// }
	// for j := 0; j < len(b); j++ {

	// 	if _, ok := tables[b[j]]; !ok {
	// 		return false
	// 	}
	// 	if j > 0 && b[j] > b[j-1] && b[j]-b[j-1] != 1 {
	// 		return false
	// 	}

	// }

	// approach 2
	var (
		idxB = -1
		j    = 0
	)
	for i := 0; i < len(b); i++ {
		if b[i] == a[0] {
			idxB = i
			break
		}
	}
	if idxB == -1 {
		return false
	}

	for i := 0; i < len(a); i++ {
		j = (idxB + i) % len(a)
		if a[i] != b[j] {
			return false
		}
	}

	return true
}
