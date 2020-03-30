package example_3

import "fmt"

const (
	TESTB = 1 << (10 * iota)
	KIB
	MiB
)

func Test()  {
	fmt.Println( TESTB)
	fmt.Println( KIB)
	fmt.Println( MiB)
}
