package main

import "fmt"

func fn(m map[int]int) {
	m = make(map[int]int)
}

func main() {
	var m map[int]int   //map is not reference variable. in go, no two variables can share the same storage loc.
	fn(m)                // they can have content which points to same shared loc but can't share same storage loc
	fmt.Println(m == nil)   //assignment to m inside fn has no effect on the value of m in main
}