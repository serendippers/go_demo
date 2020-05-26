package advanced

import "fmt"

func findTheLongestSubstring(s string) int {
	var res int
	strMap := make(map[string]bool)
	strMap["a"] = true
	strMap["e"] = true
	strMap["i"] = true
	strMap["o"] = true
	strMap["u"] = true

	for _, v := range s{
		fmt.Println(v)
	}
	fmt.Println(res)
	return 0
}