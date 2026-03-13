package main

import (
	bst "dsa/BST"
	"fmt"
	// bubblesort "dsa/BubbleSort"
	// insertionsort "dsa/InsertionSort"
	// selectionsort "dsa/SelectionSort"
	// maps "dsa/MAPS"
	// palindromestring "dsa/PalindromeString"
	// reversestring "dsa/ReverseString"
)

func main() {
	// arr := []int{8, 4, 6, 2, 7, 9, 1, 0, 4, 2, 5, 7, 9, 77, 5, 3, 4, 5, 6, 8, 3, 2, 334, 23}
	// fmt.Println("Array before sorting",arr)
	// bubblesort.BubbleSort(arr)

	// insertionsort.InsertionSort(arr)

	// selectionsort.SelectionSort(arr)
	// fmt.Println("Array after sort",arr)

	// fmt.Println(reversestring.ReverseString("Shinas Muhammed K"))

	// fmt.Println(palindromestring.PalindromeString("MalaYalAm"))
	// fmt.Println(palindromestring.PalindromeString("Shinas muhammed K k demmahuM saniHs"))
	// fmt.Println(palindromestring.PalindromeString("Shinas1; muhammed'l K k demmahuM. 1saniHs"))

	// fmt.Println(maps.CountFrequencyOfNumbers(arr))

	// fmt.Println(maps.CountFrequencyOfStrings("SHINAS MUHAMMED K"))

	// fmt.Println(maps.FirstNonRepeatingCharacter("shinasmuhammedk"))

	// dup := []int{1,2,3,4,5,4,45,4}

	// fmt.Println(maps.CheckDuplicateNumbers(dup))

	// fmt.Println(maps.TwoSum([]int{2,7,11,15}, 9))

	// fmt.Println(maps.WordFrequencyCounter("go is fun and go is fast"))

	var root *bst.TreeNode
	values := []int{4, 2, 12, 9, 15, 30, 2, 16, 7, 9, 14}
	for _, v := range values {
		root = bst.Insert(root, v)
	}
	fmt.Println("INORDER TRAVERSAL")
	bst.Inorder(root)

	fmt.Println("\nPREORDER TRAVERSAL")
	bst.Preorder(root)

	fmt.Println("\nPOSTORDER TRAVERSAL")
	bst.Postorder(root)

	node := bst.Search(root, 99)

	if node != nil {
		fmt.Println("\nValue found", node.Val)
	} else {
		fmt.Printf("\nValue not found")
	}
}
