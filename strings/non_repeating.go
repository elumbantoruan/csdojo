package strings

// NonRepeatingChar returns the first non repeating char in s
// aabcb --> c
// aaacdddbg --> c
// aabb --> ' '
func NonRepeatingChar(s string) rune {
	var kvp = make(map[rune]int)
	for _, r := range s {
		if count, ok := kvp[r]; !ok {
			kvp[r] = 1
		} else {
			count++
			kvp[r] = count
		}
	}

	for k, v := range kvp {
		if v == 1 {
			return k
		}
	}
	return ' '
}
