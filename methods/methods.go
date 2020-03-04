package methods

import (
	"math"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////
// METHODS
/////////////////////////////////////////////////////////////////////////////////////////////////////
// 1. Go does not have classes. However, you can define methods on types.
// 2. A method is a function with a special receiver argument.
// 3. The receiver appears in its own argument list between the func keyword and the method name.
type Vertex struct {
	X, Y float64
}

// func (v Vertex) abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// Remember: a method is just a function with a receiver argument.
// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// You can declare a method on non-struct types, too.
// IMPORTANT!!
// You can only declare a method with a receiver whose type is defined in the same package as the method. 
// You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
