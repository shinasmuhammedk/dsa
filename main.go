package main

import (
	// bubblesort "dsa/BubbleSort"
	// insertionsort "dsa/InsertionSort"
	// selectionsort "dsa/SelectionSort"
	palindromestring "dsa/PalindromeString"
	// reversestring "dsa/ReverseString"
	"fmt"
)

func main() {
	// arr := []int{8, 4, 6, 2, 7, 9, 1, 0, 4, 2, 5, 7, 9, 77, 5, 3, 4, 5, 6, 8, 3, 2, 334, 23}
	// fmt.Println("Array before sorting",arr)
    // bubblesort.BubbleSort(arr)
    
    // insertionsort.InsertionSort(arr)
    
    // selectionsort.SelectionSort(arr)
    // fmt.Println("Array after sort",arr)
    
    // fmt.Println(reversestring.ReverseString("Shinas Muhammed K"))
    
    fmt.Println(palindromestring.PalindromeString("MalaYalAm"))
    fmt.Println(palindromestring.PalindromeString("Shinas muhammed K k demmahuM saniHs"))
    fmt.Println(palindromestring.PalindromeString("Shinas1; muhammed'l K k demmahuM. 1saniHs"))
}