package main

import "fmt"

func main() {

	mycar := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("\nThe original array is %d\n", mycar)

	s := mycar[:5]
	printslice(s)

	s = s[1:4]
	printslice(s)

	s = s[1:2]
	printslice(s)

	s = s[:1]
	printslice(s)

	s = s[1:]
	printslice(s)
}

func printslice(s []int) {
	fmt.Printf("\n length is %d\n capacity is %d\n slice looks like %v\n\n", len(s), cap(s), s)
}
