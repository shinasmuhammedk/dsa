package maps

import "strings"

func WordFrequencyCounter(text string) map[string]int {
	freq := make(map[string]int)

	words := strings.Fields(text)
    
    for _, word := range words{
        freq[word]++
    }
    return freq
}