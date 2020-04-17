package example_7

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type BytesCounter int

type ByteWordsCounter int

type byteWriter struct {
	w io.Writer
	byteLen int64
}

func (c *BytesCounter) Write(p []byte) (int, error)  {

	*c += BytesCounter(len(p))
	return len(p), nil
}

func Exam_1()  {
	var c  BytesCounter

	_, _ = c.Write([]byte("hello"))

	fmt.Println(c)

	c = 0
	var name = "zpp"
	fmt.Fprintf(&c, "hello %s\n", name)
	fmt.Println(c)

	var wordC ByteWordsCounter
	words := "hello is a tes fu a like"
	fmt.Fprintf(&wordC, words)
	fmt.Println(wordC)
}

//练习 7.1 使用来自 ByteCounter 的思路，实现一个针对对单词和行数的计数器。你会 发现 bufio.ScanWords 非常的有用。
func (c *ByteWordsCounter) Write(p []byte) (int, error)  {
	s := strings.NewReader(string(p))
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)
	for bs.Scan()  {
		*c ++
	}
	return int(*c), nil
}


/*
写一个带有如下函数签名的函数 CountingWriter，传入一个 io.Writer 接口类型，
返回一个新的 Writer 类型把原来的 Writer 封装在里面和一个表示写入新的 Writer 字节数的 int64 类型指针
 */
func (c *byteWriter) Write(p []byte)(int, error)  {
	n, err := c.w.Write(p)
	c.byteLen += int64(len(p))
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &byteWriter{w:w, byteLen:0}
	return c, &c.byteLen
}
