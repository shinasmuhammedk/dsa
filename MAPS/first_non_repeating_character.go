package maps

func FirstNonRepeatingCharacter(s string) string {
	first := make(map[rune]int)
	for _, v := range s {
		first[v]++
	}

	for _, v := range s {
		if first[v] == 1 {
			return string(v)
		}
	}

	return ""
}
