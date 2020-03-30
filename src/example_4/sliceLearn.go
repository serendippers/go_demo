package example_4

import (
	"fmt"
)

func SliceLearn() {
	month := [...]string{1: "January", 12: "December"}
	fmt.Println(month)
	fmt.Println(len(month))
	fmt.Println(cap(month))

	arr := [10]int{}
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	ints := make([]int, 5, 8)
	fmt.Println(ints)

	var arrs [3]string = [3]string{"one", "two", "three"}

	slice1 := arrs[0:]
	slice2 := arrs[0:]

	fmt.Printf(" &slice1 %p" ,slice1)
	fmt.Printf(" &slice2 %p" ,slice2)
	slice2[1] = "test"
	fmt.Println("&slice1 == &slice2 " ,&slice1 == &slice2)
	fmt.Println("slice1 " ,slice1)
	fmt.Println("slice2 " ,slice2)

	slice1 = append(slice1, "four")
	fmt.Printf(" &slice1 %p" ,slice1)
	fmt.Printf(" &slice2 %p" ,slice2)

	fmt.Println("slice1 " ,slice1)
	fmt.Println("slice2 " ,slice2)
}
