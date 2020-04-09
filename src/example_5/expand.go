package example_5

import (
	"fmt"
	"strings"
)

/**
练习 5.9: 编写函数 expand，将 s 中的"foo"替换为 f("foo")的返回值。
strings.Replace
func Replace(s, old, new string, n int) string
返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
 */
func Expand(s string, f func(string) string) {

	fmt.Println(strings.Replace(s, "foo", f("foo"), -1))
}

func Change(s string) string {
	return strings.ToUpper(s)
}
