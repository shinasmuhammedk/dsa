package leetcode

import "sort"

func GroupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range words {

		chars := []rune(word)

		sort.Slice(chars,  func(i int, j int) bool{
            return chars[i] < chars[j]
        })
        
        key := string(chars)
        
        groups[key] = append(groups[key], word)
	}
    result := [][]string{}
    
    for _,v := range groups{
        result = append(result, v)
    }
    return result
}