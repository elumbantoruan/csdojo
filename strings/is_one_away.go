package strings

import (
	"math"
)

// IsOneAway determines (true) if there is one char different or missing in the two given strings.
// s1a = "abcde"
// s1b = "abfde"
// returns true since c is replaced by f in s1b
// s2a = "abcde"
// s2b = "abde"
// returns true since  c is removed in s2b
// s3a = "xyz"
// s3b = "xyaz"
// returns true since a is added in s3b
// s4a = "xyz"
// s4b = "xyazd"
// returns false as there is more than one away
// s5a = "xyz"
// s5b = "abc"
// returns false as there is more than one away
func IsOneAway(s1, s2 string) bool {
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}
	var diff int
	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				diff++
			}
			if diff == 2 {
				return false
			}
		}
	} else {
		i := 0
		j := 0
		in := len(s1)
		jn := len(s2)
		for i < in || j < jn {
			if s1[i] == s2[j] {
				i++
				j++
			} else {
				if diff == 1 {
					return false
				}
				if in > jn {
					i++
				} else {
					j++
				}
				diff++
			}
		}
	}
	return true
}

// IsOneAway2ndApproach determines (true) if there is one char different or missing in the two given strings.
// s1a = "abcde"
// s1b = "abfde"
// returns true since c is replaced by f in s1b
// s2a = "abcde"
// s2b = "abde"
// returns true since  c is removed in s2b
// s3a = "xyz"
// s3b = "xyaz"
// returns true since a is added in s3b
// s4a = "xyz"
// s4b = "xyazd"
// returns false as there is more than one away
// s5a = "xyz"
// s5b = "abc"
// returns false as there is more than one away
func IsOneAway2ndApproach(s1, s2 string) bool {
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}
	if len(s1) == len(s2) {
		return isOneWayCheckForSameLength(s1, s2)
	}
	if len(s1) > len(s2) {
		return isOneWayDiffOne(s1, s2)
	}
	return isOneWayDiffOne(s2, s1)

}

func isOneWayCheckForSameLength(s1, s2 string) bool {
	var diff int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}

// len(s1) = len(s2) + 1
func isOneWayDiffOne(s1, s2 string) bool {
	var diff int
	for i := 0; i < len(s2); i++ {
		if s1[i+diff] != s2[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return true
}
