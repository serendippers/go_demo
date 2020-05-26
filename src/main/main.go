package main

import (
	"advanced"
)


func main() {
	//RelocateLinkedList
	//advanced.TestRelocateLinkedList()
	//example_5.Expand("fooaaaaaaoofsfoaofooaaaajj", example_5.Change)

	//uglyNumber := advanced.UglyNumber1(8)
	//fmt.Printf("UglyNumber1 is %d\n", uglyNumber)

	//fmt.Println(math.Hypot(3, 4))

	//example_6.TestDistance()
	//example_7.Exam_1()
	//example_7.Test_7_4()
	//example_7.Test_7_5()
	//example_7.TestExample_7_6()
	//example_8.TestExample_8_1()
	//example_8.TestExample_8_2()
	//example_8.TestExample_8_4()
	//example_8.TestExample_8_4_2()
	//example_8.TestExample_8_4_4()

	//var wg sync.WaitGroup
	//var x, y int
	//go func() {
	//	x = 1                   // A1
	//	fmt.Print("y:", y, " ") // A2
	//	wg.Done()
	//}()
	//go func() {
	//	y = 1                   // B1
	//	fmt.Print("x:", x, " ") // B2
	//	wg.Done()
	//}()
	//wg.Wait()

	cache := advanced.Constructor(2)

	cache.Put(1, 1);
	cache.Put(2, 2);
	cache.Get(1);       // 返回  1
	cache.Put(3, 3);    // 该操作会使得关键字 2 作废
	cache.Get(2);       // 返回 -1 (未找到)
	cache.Put(4, 4);    // 该操作会使得关键字 1 作废
	cache.Get(1);       // 返回 -1 (未找到)
	cache.Get(3);       // 返回  3
	cache.Get(4);       // 返回  4

}
