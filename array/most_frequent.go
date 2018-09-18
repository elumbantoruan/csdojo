package array

import (
	"math"
)

// MostFrequent returns item of an array that appears most frequently
func MostFrequent(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var items = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := items[nums[i]]; !ok {
			items[nums[i]] = 1
		} else {
			items[nums[i]] = val + 1
		}
	}
	var (
		mf  int
		max = math.MinInt32
	)
	for k, v := range items {
		if v > max {
			max = v
			mf = k
		}
	}
	return mf
}
