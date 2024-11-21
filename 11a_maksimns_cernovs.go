package main

import (
	"fmt"
	"os"
)

func main() {
	var n int //9
	var s int
	var d int

	fmt.Print("n: ")
	fmt.Fscan(os.Stdin, &n)

	fmt.Print("s: ")
	fmt.Fscan(os.Stdin, &s)

	fmt.Print("d: ")
	fmt.Fscan(os.Stdin, &d)

	var r int //sarkans tomats
	r = 1
	// var i int
	// if s < n {
	// 	i = r + 2 * d
	// 	s
	// }
	// fmt.Println(i)
	var res int
	var p1 int
	var p2 int
	if s <= n {
		if (s-d) > 0 && (s+d) < n {
			p1 = s + d
			p2 = s - d
			res = p1 - p2 + r
		}
	}
	fmt.Println(res)

}
