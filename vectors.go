//Package vectors is for working with vectors of N dimensions and operate with theme.
//to define a vector you shold create theme with the New method, all the operations
//are need or outputs a Vector object
package vectors

import (
	"errors"
	"math"
)

//Vector is the representation of a vector of N dimensions, it has the attributes:
//Components: A float64 slice (to define the components of the wanted vector)
//Magnitude: A float64 (the magnitude of the vector)
type Vector struct {
	Components []float64
	Magnitude  float64
}

// New creates a new Vector object. Every time you call this function the
//Magnitude attribute is calculated
func New(d []float64) *Vector {
	v := new(Vector)
	v.Components = d
	v.Magnitude = v.GetMagnitude()
	return v
}

//GetMagnitude return the magnitude of a vector
func (v *Vector) GetMagnitude() float64 {

	var result float64

	if len(v.Components) == 0 {
		return 0
	}

	for i := 0; i < len(v.Components); i++ {
		result += v.Components[i] * v.Components[i]
	}

	return math.Sqrt(result)
}

//Ones make a Vector filled with ones
func Ones(n int) *Vector {
	var com []float64
	if n != 0 {
		for i := 0; i < n; i++ {
			com = append(com, 1)
		}
		return New(com)
	}
	return nil
}

//Add returns A + B where A and B are Vectors
func Add(A, B *Vector) *Vector {

	if !Checklenght(A, B) {
		return nil
	}

	var sum []float64

	for i := 0; i < len(A.Components); i++ {
		sum = append(sum, A.Components[i]+B.Components[i])
	}

	return New(sum)
}

//Substract returns A + (-B) or A - B where A and B are Vectors
func Substract(A, B *Vector) *Vector {
	return Add(A, Minus(B))
}

//ScalarProduct returns the Vector scaled with a factor of a (a*A)
func ScalarProduct(a float64, A *Vector) *Vector {
	var newA []float64

	for i := 0; i < len(A.Components); i++ {
		newA = append(newA, a*A.Components[i])
	}

	return New(newA)
}

//Minus returns the Vector -A
func Minus(A *Vector) *Vector {
	return ScalarProduct(-1, A)
}

//DotProduct computes the point product of the vectors A and B. Returns a float64
func DotProduct(A, B *Vector) (float64, error) {

	var result float64

	if !Checklenght(A, B) {
		return 0, errors.New("The input vectors are not the same lenght")
	}

	for i := 0; i < len(A.Components); i++ {
		result += A.Components[i] * B.Components[i]
	}
	return result, nil

}

//Angle returns the Angle between Vectors A and B
func Angle(A, B *Vector) (float64, error) {

	var dot, err = DotProduct(A, B)

	if err != nil {
		return 0, err
	}
	return math.Acos(dot/(A.Magnitude*B.Magnitude)) * (180 / math.Pi), err
}

// Checklenght check if two Vectors have the same dimensions (number of Components)
func Checklenght(A, B *Vector) bool {

	if len(A.Components) == len(B.Components) {
		return true
	}
	return false

}
