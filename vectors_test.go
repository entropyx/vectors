package vectors

import (
	"math"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TwoVectors(f func(A, B *Vector)) func() {
	return func() {

		A := New([]float64{4, 3})
		B := New([]float64{1, 1})

		f(A, B)

	}

}

func TestDotProduct(t *testing.T) {
	Convey("Testing dot product", t, TwoVectors(func(A *Vector, B *Vector) {
		got, _ := DotProduct(A, B)
		So(got, ShouldEqual, 7)

		A.Components = []float64{6, 2, 3}
		B.Components = []float64{6, 2, 3}
		got, _ = DotProduct(A, B)
		So(got, ShouldEqual, 49)

	}))
}

func TestSumm(t *testing.T) {
	Convey("Testing Sum of vectors", t, TwoVectors(func(A *Vector, B *Vector) {
		got := Add(A, B)
		So(got, ShouldResemble, New([]float64{5, 4}))
	}))
}

func TestGetMagnitude(t *testing.T) {

	var tests = []struct {
		input *Vector
		want  float64
	}{
		{New([]float64{4, 3}), 5},
		{New([]float64{}), 0},
	}

	for _, test := range tests {
		Convey("Checking the Magnitude of a vector", t, func() {
			So(test.input.GetMagnitude(), ShouldEqual, test.want)

		})
	}

}

//
// 	if  != test.want {
// 		t.Errorf("GetMagnitude() of the vector %v  mismatch %v", test.input.Components, test.want)
// 	}
// }

func TestOnes(t *testing.T) {

	var tests = []struct {
		input int
		want  *Vector
	}{
		{1, New([]float64{1})},
		{0, nil},
	}

	for _, test := range tests {
		if got := Ones(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Ones(%v) = %v /n Shoul be %v ", test.input, got, test.want)
		}
	}
}

func TestAdd(t *testing.T) {

	var tests = []struct {
		A, B *Vector
		want *Vector
	}{
		{New([]float64{1, 2}), New([]float64{2, 1, 4}), nil},
		{New([]float64{1, 2}), New([]float64{2, 1}), New([]float64{3, 3})},
		{New([]float64{1.3, -2}), New([]float64{2, 2}), New([]float64{3.3, 0})},
		{New([]float64{1.3, math.Sqrt(2)}), New([]float64{math.Sqrt(3), 2}), New([]float64{1.3 + math.Sqrt(3), 2 + math.Sqrt(2)})},
	}

	for _, test := range tests {
		if got := Add(test.A, test.B); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Add(%v,%v) = %v /n Shoul be %v ", test.A, test.B, got, test.want)
		}
	}

}

func TestSubstract(t *testing.T) {

	var tests = []struct {
		A, B *Vector
		want *Vector
	}{
		{New([]float64{1, 2}), New([]float64{2, 1, 4}), nil},
		{New([]float64{1, 2}), New([]float64{2, 1}), New([]float64{-1, 1})},
		{New([]float64{1.3, -2}), New([]float64{2, 2}), New([]float64{-0.7, -4})},
		{New([]float64{1.3, math.Sqrt(2)}), New([]float64{math.Sqrt(3), 2}), New([]float64{1.3 - math.Sqrt(3), -2 + math.Sqrt(2)})},
	}

	for _, test := range tests {
		if got := Substract(test.A, test.B); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Substract(%v,%v) = %v /n Shoul be %v ", test.A, test.B, got, test.want)
		}
	}

}

// func TestDotProduct(t *testing.T) {
//
// 	var tests = []struct {
// 		A, B *Vector
// 		want float64
// 		err  error
// 	}{
// 		{New([]float64{1, 2}), New([]float64{2, 1, 4}), 0, nil},
// 		{New([]float64{1, 2}), New([]float64{2, 1}), 4, nil},
// 		{New([]float64{1.3, -2}), New([]float64{2, 2}), -1.4, nil},
// 		{New([]float64{1.3, math.Sqrt(2)}), New([]float64{math.Sqrt(3), 2}), (1.3*math.Sqrt(3) + 2*math.Sqrt(2)), nil},
// 	}
//
// 	for _, test := range tests {
// 		if got, _ := DotProduct(test.A, test.B); got != test.want {
// 			t.Errorf("DotProduct(%v,%v) = %v /n Shoul be %v ", test.A, test.B, got, test.want)
// 		}
// 	}
//
// }

func TestAngle(t *testing.T) {

	var tests = []struct {
		A, B *Vector
		want float64
		err  error
	}{
		{New([]float64{1, 2}), New([]float64{2, 1, 4}), 0, nil},
		{New([]float64{1, 2}), New([]float64{-2, 1}), 90, nil},
		{New([]float64{0, 2}), New([]float64{0, -2}), 180, nil},
	}

	for _, test := range tests {
		if got, _ := Angle(test.A, test.B); got != test.want {
			t.Errorf("Angle(%v,%v) = %v /n Shoul be %v ", test.A, test.B, got, test.want)
		}
	}

}
