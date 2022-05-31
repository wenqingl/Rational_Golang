package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Floater64 interface {
	// Converts a value to an equivalent float64.
	toFloat64() float64
}

type Rationalizer interface {

	// 5. Rationalizers implement the standard Stringer interface.
	fmt.Stringer

	// 6. Rationalizers implement the Floater64 interface.
	Floater64

	// 2. Returns the numerator.
	Numerator() int

	// 3. Returns the denominator.
	Denominator() int

	// 4. Returns the numerator, denominator.
	Split() (int, int)

	// 7. Returns true iff this value equals other.
	Equal(other Rationalizer) bool

	// 8. Returns true iff this value is less than other.
	LessThan(other Rationalizer) bool

	// 9. Returns true iff the value equal an integer.
	IsInt() bool

	// 10. Returns the sum of this value with other.
	Add(other Rationalizer) Rationalizer

	// 11. Returns the product of this value with other.
	Multiply(other Rationalizer) Rationalizer

	// 12. Returns the quotient of this value with other. The error is nil
	// if its is successful, and a non-nil if it cannot be divided.
	Divide(other Rationalizer) (Rationalizer, error)

	// 13. Returns the reciprocal. The error is nil if it is successful,
	// and non-nil if it cannot be inverted.
	Invert() (Rationalizer, error)

	// 14. Returns an equal value in lowest terms.
	ToLowestTerms() Rationalizer
} // Rationalizer interface

type Rational struct {
	numerator   int
	denominator int
}

// 2.
func (r Rational) Numerator() int {
	return r.numerator
}

// 3.
func (r Rational) Denominator() int {
	return r.denominator
}

// 4.
func (r Rational) Split() (int, int) {
	return r.numerator, r.denominator
}

// 5.
func (r Rational) String() string {
	return fmt.Sprintf("%v/%v", r.numerator, r.denominator)
}

// 6.
func (r Rational) toFloat64() float64 {
	return float64(r.numerator) / float64(r.denominator)
}

// 7.
func (r Rational) Equal(other Rationalizer) bool {
	r_gdc := GCD(r.numerator, r.denominator)
	other_gdc := GCD(other.Numerator(), other.Denominator())

	if r.denominator/r_gdc == other.Denominator()/other_gdc && r.numerator/r_gdc == other.Numerator()/other_gdc {
		return true
	}
	return false

}

// GCD
func GCD(m, n int) int {
	for n != 0 {
		t := n
		n = m % n
		m = t
	}
	return m
}

// 8.
func (r Rational) LessThan(other Rationalizer) bool {
	if r.toFloat64() < other.toFloat64() {
		return true
	} else {
		return false
	}
}

// 9.
func (r Rational) IsInt() bool {
	if r.numerator%r.denominator == 0 {
		return true
	} else {
		return false
	}
}

// 10.
func (r Rational) Add(other Rationalizer) Rationalizer {
	a := r.numerator*other.Denominator() + other.Numerator()*r.denominator
	b := r.denominator * other.Denominator()

	gcd := GCD(a, b)
	return Rational{a / gcd, b / gcd}
}

// 11.
func (r Rational) Multiply(other Rationalizer) Rationalizer {
	a := r.numerator * other.Numerator()
	b := r.denominator * other.Denominator()

	gcd := GCD(a, b)
	return Rational{a / gcd, b / gcd}
}

// 12.
func (r Rational) Divide(other Rationalizer) (Rationalizer, error) {
	a := r.numerator * other.Denominator()
	b := r.denominator * other.Numerator()

	if b == 0 {
		return Rational{0, 0}, errors.New("can not divided by zero")
	} else {
		gcd := GCD(a, b)
		return Rational{a / gcd, b / gcd}, nil
	}
}

// 13.
func (r Rational) Invert() (Rationalizer, error) {
	a := r.numerator
	b := r.denominator

	if a == 0 {
		return Rational{0, 0}, errors.New("denominator cannot be zero")
	} else {
		return Rational{b, a}, nil
	}
}

// 14.
func (r Rational) ToLowestTerms() Rationalizer {
	gcd := GCD(r.numerator, r.denominator)
	return Rational{r.numerator / gcd, r.denominator / gcd}
}

// 15. Harmonic sum
func HarmonicSum(n int) Rationalizer {
	var sum Rationalizer
	sum = Rational{1, 1}

	for i := 2; i <= n; i++ {
		sum = sum.Add(Rational{1, i})
	}
	return sum
}

// insertion sort for int
func insertionSortInt(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j], a[j-1] = a[j-1], a[j] // swap a[j], a[j-1]
		}
	}
	return a
}

// insertion sort for
func insertionSortString(a []string) []string {
	n := len(a)
	if n < 2 {
		return a
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j], a[j-1] = a[j-1], a[j] // swap a[j], a[j-1]
		}
	}
	return a
}

func insertionSortRational(a []Rationalizer) []Rationalizer {
	n := len(a)
	if n < 2 {
		return a
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && a[j].LessThan(a[j-1]); j-- {
			a[j], a[j-1] = a[j-1], a[j] // swap a[j], a[j-1]
		}
	}
	return a
}

// random string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	// average for different n
	average_int := make([]float64, 10)
	average_str := make([]float64, 10)
	average_rat := make([]float64, 10)

	// run for size 1000 - 10000
	for i := 0; i < 10; i++ {
		n := 1000 * (i + 1)

		// record the sum of time for 3 type
		sum := make([]int64, 3)

		for j := 0; j < 3; j++ {
			// ----------------- integer type -----------------
			// create integer list
			IntList := make([]int, n)
			for m := 0; m < n; m++ {
				IntList[m] = rand.Intn(10000-(-10000)) + (-10000)
			}

			// record the runtime of integer
			start := time.Now() // record the start time
			insertionSortInt(IntList)
			end := time.Now()                        // record the end time
			elapsed := end.Sub(start).Microseconds() // runtime
			sum[0] += elapsed

			// ----------------- string type -----------------
			// create string list
			StrList := make([]string, n)
			for m := 0; m < n; m++ {
				StrList[m] = randStr(4)
			}

			// record the runtime of string
			start = time.Now()
			insertionSortString(StrList)
			end = time.Now()
			elapsed = end.Sub(start).Microseconds()
			sum[1] += elapsed

			// ----------------- Rational type -----------------
			// create rational list
			RatList := make([]Rationalizer, n)
			for m := 0; m < n; m++ {
				numerator := rand.Intn(10000-(-10000)) + (-10000)
				denominator := rand.Intn(10000-(-10000)) + (-10000)

				// check valid rational
				if denominator == 0 {
					m--
				} else {
					RatList[m] = Rational{numerator, denominator}
				}
			}

			// record the runtime of rational
			start = time.Now()
			insertionSortRational(RatList)
			end = time.Now()
			elapsed = end.Sub(start).Microseconds()
			sum[2] += elapsed
		}

		average_int[i] = float64(sum[0]) / 3
		average_str[i] = float64(sum[1]) / 3
		average_rat[i] = float64(sum[2]) / 3
	}

	// print the output
	fmt.Println("runtime of integer type:")
	for i := 0; i < 10; i++ {
		//fmt.Println("n =", 1000*(i+1), ":", average_int[i], "microseconds")
		fmt.Printf("n = %v: %.2f microseconds\n", 1000*(i+1), average_int[i])
	}

	fmt.Println("\nruntime of string type:")
	for i := 0; i < 10; i++ {
		//fmt.Println("n =", 1000*(i+1), ":", average_str[i], "microseconds")
		fmt.Printf("n = %v: %.2f microseconds\n", 1000*(i+1), average_str[i])
	}

	fmt.Println("\nruntime of rational type:")
	for i := 0; i < 10; i++ {
		//fmt.Println("n =", 1000*(i+1), ":", average_rat[i], "microseconds")
		fmt.Printf("n = %v: %.2f microseconds\n", 1000*(i+1), average_rat[i])
	}
}
