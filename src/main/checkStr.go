package main

func check(s1, s2 string) bool {
	if len(s1)!=len(s2) {
		return false
	}

	byteMap := make(map[byte]int)

	for _, s := range s2 {
		_, ok := byteMap[byte(s)]
		if ok {
			byteMap[byte(s)] ++
		} else {
			byteMap[byte(s)] = 1
		}
	}

	for _, s := range s1 {
		n, ok := byteMap[byte(s)]
		if ok && n>0 {
			byteMap[byte(s)] --
		} else {
			return false
		}
	}
	return true
}
