package array

// CommonElements returns the common items of the two arrays
// A = {1,3,4,6,7,9}
// B = {1,2,4,5,9,10}
func CommonElements(nums1 []int, nums2 []int) []int {
	var commons []int

	var (
		i = 0
		j = 0
	)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			commons = append(commons, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}

	return commons
}
