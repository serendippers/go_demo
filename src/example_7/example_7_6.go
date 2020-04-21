package example_7

import (
	"fmt"
	"sort"
)

func TestExample_7_6 () {
	//learnSort()

	a := []int{3,2,1,2,3}
	fmt.Println(IsPalindrome(sort.IntSlice(a)))
	b := []int{3,2,1,3,2,3}
	fmt.Println(IsPalindrome(sort.IntSlice(b)))
}

func learnSort()  {

	values := []int{3,1,4,1}

	fmt.Println(sort.IntsAreSorted(values))
	fmt.Println(values)

	sort.Ints(values)
	fmt.Println(sort.IntsAreSorted(values))
	fmt.Println(values)
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(sort.IntsAreSorted(values))
	fmt.Println(values)
}

/*
练习 7.10: sort.Interface 类型也可以适用在其它地方。
编写一个 IsPalindrome(s sort.Interface) bool 函数表明序列 s 是否是回文序列，换句话说反向排序不会改变这 个序列。
假设如果!s.Less(i, j) && !s.Less(j, i)则索引 i 和 j 上的元素相等。
 */
func IsPalindrome(s sort.Interface) bool  {

	if s.Len() < 1 {
		return false
	}

	for i, j := 0, s.Len()-1; i<j; {
		if !s.Less(i, j) && !s.Less(j, i){
			i ++
			j --
		} else {
			return false
		}
	}
	return true
}