package main

func inlined() []int {
	var a [5]int
	return a[:]
}

//go:noinline
func no_inline() []int {
	var b [5]int
	return b[:]
}

func main() {
	var local_array [5]int
	var local_var int       // on running the main program, we can see
	println(no_inline())    // address of no_inline.b will be different as its stored in heap (due to noinline)
	println(inlined())      // whereas inlined.a is stored in stack seen in result of  "go build -gcflags='-m' noninline.go"
	println(local_array[:]) // that compiler realises that a does not escape beyond its scope and can be placed on stack
	println(&local_var)     // go:noinline is the Go pragma used here to demonstrate no inlining of the small functions in main function
}


