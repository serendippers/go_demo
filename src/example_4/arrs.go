package example_4

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD int = iota
	EUR
	GBP
	RMB
)

func CurrencyTest()  {
	symbol := [...]string{USD:"$", GBP:"??", RMB:"¥"}
	fmt.Println(RMB, symbol[RMB])
	fmt.Println(symbol)
	fmt.Println(len(symbol))
	fmt.Println(cap(symbol))
}

func ArrTest() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

}

func ShaTest()  {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	/*
	%x 指定以十六进制的格 式打印数组或 slice 全部的元素
	%t 副词参数是用于打印布尔型数据
	%T 副词参数是 用于显示一个值对应的数据类型
	*/
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println( c1[0] >> uint(1) != c2[0]>>uint(1))

}