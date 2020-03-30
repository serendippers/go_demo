package main

import (
	"fmt"
)

func main() {
	//example_3
	/*s := comma("+123456789")
	fmt.Println(s)

	s1 := "qwertyuqqv"
	s2 := "uytrewqqqb"

	t := check(s1, s2)
	fmt.Println(t)

	fmt.Println(formatToInt("123"))

	var v example_3.Flags = example_3.FlagMulticast | example_3.FlagUp
	example_3.IsUp(v)

	example_3.Test()*/

	//example_4.ArrTest()
	//example_4.ShaTest()
	//example_4.SliceLearn()
	//example_4.CurrencyTest()
	//a := [...]int{0, 1, 2, 3, 4, 5}
	//example_4.Reverse(a[:])
	//could not launch process: debugserver or lldb-server not found: install XCode's command line tools or lldb-server
	//fmt.Println(a)

	nums := []int{11, 7, 2, 15}
	target := 9
	res := twoSumHash(nums, target)

	fmt.Println(res)
}



func stack(){

}


func twoSumHash(nums []int, target int) [2]int {

	res := [2]int{}
	if len(nums) == 0 {
		return res
	}


	numMap := make(map[int]int)

	for i, j := range nums {
		numMap[j] = i
	}
	for i, j := range nums {
		temp := target - j

		_,ok := numMap[temp]
		if ok {
			res[0] = i;
			res[1] = numMap[temp]
			return res
		}
	}
	return res
}



