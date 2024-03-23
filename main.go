package main

import "ttharsh.gobook/ch4/xkcd"

func main() {
	// // ch1.Solve()
	// ch1.Solve12()
	// endTime := time.Now()
	// fmt.Println("Execution time: " , endTime.Sub(startTime))
	// ch1.Dup2();
	// ch1.Lissajous(os.Stdout)
	// ch1.FetchAll()
	// ch1.InitGIFServer()
	// fmt.Printf("%s", tempconv.CToF(tempconv.AbsoluteZeroKelvin))
	// startTime := time.Now()
	// popcount.PopCountv2(100)
	// endTime := time.Now()
	// fmt.Printf("Execution time: %d", endTime.Sub(startTime))
	// var a int = 077
	// fmt.Printf("%d", a)
	// fmt.Printf("%s\n", ch3.CommaRecursive("12578"))
	// fmt.Printf("%s\n", ch3.CommaBytesFloat("-121256.123"))
	// fmt.Printf("%t\n", ch3.Anagram("hello", "lllhe"))
	// shaBitDiff := ch4.ShaDiff("hell", "hello")
	// fmt.Printf("Bit Difference between SHA256: %d\n", shaBitDiff)

	// Ex 4.5
	// arr := []string{"h", "h", "a", "h", "a", "a", "", "x"}
	// fmt.Printf("Current slice: %v\n", arr)
	// arr = ch4.RemoveDups(arr)
	// fmt.Printf("Slice after adjacent duplicate removal: %v\n", arr)

	// Ex 4.6
	// arr := []byte{'a', 'a', '\n', '\t', ' ', 'x'}
	// fmt.Printf("Current slice: %v\n", arr)
	// arr = ch4.SquashSpace(arr)
	// fmt.Printf("Slice after squashing space: %v\n", arr)

	// file, err := os.Open("big.txt")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer file.Close()

	// ch4.WordCount(file)
	xkcd.StartApp()
}
