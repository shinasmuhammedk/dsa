package maps

func CountFrequencyOfStrings(s string) map[string]int{
    count := make(map[string]int)
    for _,v := range s{
        count[string(v)]++
    }
    return count
}