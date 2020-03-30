package main

import (
	"bytes"
	"strings"
)

func comma(s string) string {

	var buf bytes.Buffer
	i:=0
	if strings.HasPrefix(s,"+") || strings.HasPrefix(s, "-") {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	for  n := len(s); n>0; {
		m := n% 3
		if m == 0 {
			buf.WriteString(s[i:i+3])
			n -= 3
			i+=3
		} else {
			buf.WriteString(s[:m])
			n -= m
			i+=m
		}
		if(n> 0) {
			buf.WriteString(",")
		}
	}
	return buf.String()
}