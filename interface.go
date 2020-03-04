package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/google/go-cmp/cmp"
	"testing.com/interface/concurrency"
	. "testing.com/interface/methods"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS
/////////////////////////////////////////////////////////////////////////////////////////////////////

// In Go, the variable name comes before the variable type (x int, y int) as well as the return type
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add_2(x, y int) int {
	return x + y
}

// A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// A return statement without arguments returns the named return values.
	return
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(add(1, 2))
	fmt.Println(add_2(2, 3))
	last_name, first_name := swap("aldrich", "choi")
	fmt.Println(last_name, first_name)
	fmt.Println(split(17))
	// Variables that starts with a small letter is not exportable
	// fmt.Print(math.pi)
	fmt.Println(math.Pi)

	/////////////////////////////////////////////////////////////////////////////////////////////////////
	// VARIABLES
	/////////////////////////////////////////////////////////////////////////////////////////////////////
	// The var statement declares a list of variables; as in function argument lists, the type is last.
	// var i int
	// A var declaration can include initializers, one per variable.
	var j, k int = 1, 2
	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var c, python, java = true, false, "no!"
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	i := 3
	fmt.Println(i, j, k, c, python, java)
	// Variables declared without an explicit initial value are given their zero value.
	var z int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", z, f, b, s)
	// The expression T(v) converts the value v to the type T.
	// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
	i2 := 42
	f2 := float64(i2)
	var u uint = uint(f)
	fmt.Println(i2, f2, u)
	// Constants are declared like variables, but with the const keyword.
	const Birthday = "12/05"
	fmt.Println(Birthday)

	/////////////////////////////////////////////////////////////////////////////////////////////////////
	// FLOW CONTROL
	/////////////////////////////////////////////////////////////////////////////////////////////////////
	
	// LOOPS
	// Go has only one looping construct, the for loop.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("The sum is", sum)

	// The init and post statements are optional.
	sum2 := 1
	for ; sum2 <= 100; {
		sum2 += sum2
	}
	fmt.Println("The sum2 is", sum2)

	// you can drop the semicolons: C's while is spelled for in Go.
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println("The sum3 is", sum3)

	// Infinite loop
	// for {
	// }

	// IF...ELSE
	// SWITCH
	// DEFER



	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// CONCURRENCY
	/////////////////////////////////////////////////////////////////////////////////////////////////////
	// A goroutine is a lightweight thread managed by the Go runtime.
	go concurrency.Say("hello world!")
	s2 := []int{7, 2, 8, -9, 4, 0}
	// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
	// Channels must be created before use:
	c2 := make(chan int)
	go concurrency.Sum(s2[:len(s2)/2], c2)
	go concurrency.Sum(s2[len(s2)/2:], c2)
	x, y := <-c2, <-c2 // receive from c

	fmt.Println(x, y, x+y)
	// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
	ch := make(chan int, 2)
	// ch := make(chan int, 1)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Print("Abs is ")
	fmt.Println(v.Abs())
	// fmt.Println(Abs(v))

	myfloat := MyFloat(1)
	fmt.Println(myfloat.Abs())
	
}
