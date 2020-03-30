package main

import "strconv"

func formatToInt(s string) int64 {

	/*i, err := strconv.Atoi(s)
	if err!=nil {

	}*/
	i, _ := strconv.ParseInt(s, 10, 0)
	return i
}
