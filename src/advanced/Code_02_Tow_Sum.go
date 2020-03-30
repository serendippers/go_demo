package advanced

/**
给定一个数组arr，和一个整数aim，请返回哪两个位置的数可 以加出aim来。
例如
arr = {2, 7, 11, 15}，target = 9
返回{0,1}，因为arr[0] + arr[1] = 2 + 7 = 9 可以假设每个数组里只有一组答案

解法一：使用hash
 */
func twoSumHash(nums []int, target int) [2]int {

	res := [2]int{}
	if len(nums) == 0 {
		return res
	}


	numMap := make(map[int]int)

	//存入hash
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

//解法二 双指针
func twoSum(nums []int, target int)  {

}
