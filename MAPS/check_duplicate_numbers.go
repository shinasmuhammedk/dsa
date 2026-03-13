package maps

func CheckDuplicateNumbers(arr []int) bool {
	seen := make(map[int]bool)

	for _, v := range arr {
		if seen[v]{
            return true
        }
        seen[v] = true
	}
    return false
}
