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
	res := twoSum(nums, target)

	fmt.Println(res)
}



func stack(){

}


func twoSum(nums []int, target int) [2]int {
	indices := []int{}

	for i, _ := range nums {
		indices = append(indices, i)
	}

	sort(nums, indices)
	left := 0;
	right := len(nums)-1
	sum := 0
	for left< right {
		sum = nums[left] + nums[right]
		if sum == target {
			return [2]int{indices[left], indices[right]}
		}
		if sum < target {
			left ++
		}
		if sum > target {
			right --
		}
	}
	return [2]int{}
}

func sort(nums, indices []int) {

	//构建大顶堆
	for i, _ := range nums {
		heapInsert(nums, indices, i)
	}

	//排序
	for i := len(nums) - 1; i > 0; i-- {
		swap(nums, indices, 0, i)
		heapify(nums, indices, i)
	}

}

/**
构建堆
*/
func heapInsert(nums, indices []int, i int) {
	for i > 0 {
		p := (i - 1) / 2
		if nums[p] >= nums[i] {
			return
		}
		swap(nums, indices, i, p)
		i = p
	}
}

/**
调整堆结构，使之恢复大顶堆
*/
func heapify(nums, indices []int, size int) {
	i := 0
	left := 1
	right := 2
	largest := i
	for left < size {
		if nums[left] > nums[i] {
			largest = left
		}
		if nums[right] > nums[i] {
			largest = right
		}
		if largest == i {
			break
		}
		swap(nums, indices, largest, i)
		i = largest
		left = largest*2 + 1
		right = largest*2 + 2
	}

}

/*
交换位置
*/
func swap(nums, indices []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp

	temp = indices[i]
	indices[i] = indices[j]
	indices[j] = temp

}



