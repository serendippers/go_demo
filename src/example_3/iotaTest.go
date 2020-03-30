package example_3

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1<<iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast


)

func IsUp(f Flags) bool  {
	fmt.Println(f&FlagUp == FlagUp)
	fmt.Println( FlagUp)
	fmt.Println( f)
	fmt.Println( f&FlagUp)

	return f&FlagUp == FlagUp
}
