package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var packageLevelVar = "I'm outside function"

// shortHandDeclaration := "I'm outside function" // invalid

func main() {
	fmt.Println(packageLevelVar)

	var str string
	var strWithInitialization string = "I'm initialized string"
	var str1, str2 string
	fmt.Println(str, strWithInitialization, str1, str2)

	var strWithoutExplicitType = "I'm string data type"
	strShorthandDeclaration := "I'm string with shorthand declaration" // inside function only
	fmt.Println(strWithoutExplicitType, strShorthandDeclaration)

	const myStr = "I'm string constant" // seq of bytes
	fmt.Println("myStr[0]=", myStr[0])  // 73, ascii code
	// myStr[0] = "x" // invalid
	const myStrWithSmile = "I'm 😀"
	var multiLineStr string = `I'm 
multiline 
string`

	const a1 = 10
	const a2 = a1

	var b1 = 10
	fmt.Println(b1)
	// const b2 = b1 // invalid

	var myBool1 = rand.Int() > 10
	var myBool2 = rand.Int() < 10
	var iAmTrueOrFalse = myBool1 || myBool2
	_, _ = multiLineStr, iAmTrueOrFalse

	// Logical	&& || !
	// Equality	== !=

	var i int = 404                     // Platform dependent
	var i8 int8 = 127                   // -128 to 127
	var i16 int16 = 32767               // -2^15 to 2^15 - 1
	var i32 int32 = -2147483647         // -2^31 to 2^31 - 1
	var i64 int64 = 9223372036854775807 // -2^63 to 2^63 - 1
	_, _, _, _, _ = i, i8, i16, i32, i64

	var ui uint = 404                     // Platform dependent
	var ui8 uint8 = 255                   // 0 to 255
	var ui16 uint16 = 65535               // 0 to 2^16
	var ui32 uint32 = 2147483647          // 0 to 2^32
	var ui64 uint64 = 9223372036854775807 // 0 to 2^64
	var uiptr uintptr                     // Integer value of a memory address
	_, _, _, _, _, _ = ui, ui8, ui16, ui32, ui64, uiptr

	//type byte = uint8
	//type rune = int32

	var singleByte byte = 'b'
	var singleRune rune = '😀' // unicode code point
	_, _ = singleByte, singleRune

	var f32 float32 = 1.7812 // IEEE-754 32-bit
	var f64 float64 = 3.1415 // IEEE-754 64-bit
	_, _ = f32, f64

	// Arithmetic	+ - * / %
	// Comparison	== != < > <= >=
	// Bitwise	& | ^ << >>
	// Increment/Decrement	++ --
	// Assignment	= += -= *= /= %= <<= >>= &= |= ^=

	var iZero int     // 0
	var fZero float64 // 0
	var bZero bool    // false
	var sZero string  // ""
	_, _, _, _ = iZero, fZero, bZero, sZero

	mainQuestion := 42
	f := float64(mainQuestion)
	u := uint(f)
	_ = u

	iAmInt64, _ := strconv.ParseInt("100500", 10, 64)
	iAmInt64Overflow, err := strconv.ParseInt("100500", 10, 8)
	fmt.Println(err)                                                                     // strconv.ParseInt: parsing "100500": value out of range
	fmt.Printf("%v %T, %v %T\n", iAmInt64, iAmInt64, iAmInt64Overflow, iAmInt64Overflow) // 100500 int64, 127 int64

	x := rand.Int()
	if x > 5 {
		fmt.Println("x is gt 5")
	} else if x > 10 {
		fmt.Println("x is gt 10")
	} else {
		fmt.Println("else case")
	}

	if x := rand.Int(); x > 5 {
		fmt.Println("x is gt 5")
	}

	day := getCurrentDayName()
	switch day.String() {
	case "Monday":
		fmt.Println("I like my job!")
	case "Friday":
		fmt.Println("Deployment freeze!")
	default:
		fmt.Println("Focused coding!")
	}

	switch day := getCurrentDayName(); day.String() {
	case "Monday":
		fmt.Println("I like my job!")
	case "Friday":
		fmt.Println("Deployment freeze!")
	default:
		fmt.Println("Focused coding!")
	}

	switch day := getCurrentDayName(); day.String() {
	case "Monday":
		fmt.Println("I like my job!")
		fallthrough
	case "Tuesday":
		fmt.Println("I like my job!")
	case "Friday":
		fmt.Println("Deployment freeze!")
	default:
		fmt.Println("Focused coding!")
	}

	switch {
	case x > 5:
		fmt.Println("x > 5")
	default:
		fmt.Println("x <= 5")
	}

	// init; condition; post
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		x := rand.Int()
		if x > 100 {
			break
		}
		if x > 50 {
			continue
		}
		fmt.Println(x)
	}

	myFunction()
	myFunctionWithParam("hi")
	myFunctionWithMultiParams("hi", "go")
	_ = myFunctionSingleReturn("test")
	myFunctionFirstClass()
	myFunctionAnon()
	fClosure := myFunctionWithClosure()
	fClosure(10)
	fmt.Println(fClosure(20)) // 30

	add(1, 2, 3, 4, 5)
}

func myFunction() {}

func myFunctionWithParam(p1 string) {
	fmt.Println(p1)
}

func myFunctionWithMultiParams(p1, p2 string) {
	fmt.Println(p1, p2)
}

func myFunctionSingleReturn(p string) string {
	return p
}

func myFunctionMultiReturn(p string) (string, int) {
	msg := fmt.Sprintf("%s function", p)
	return msg, 10
}

func myFunctionNamedReturn(p1 string) (s string, i int) {
	s = fmt.Sprintf("%s function", p1)
	i = 10

	return
}

func myFunctionFirstClass() {
	fn := func() {
		fmt.Println("inside fn")
	}

	fn()
}

func myFunctionAnon() {
	func() {
		fmt.Println("inside fn")
	}()
}

func myFunctionWithClosure() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v

		return sum
	}
}

func add(values ...int) int {
	sum := 0

	for _, v := range values {
		sum += v
	}

	return sum
}

func getCurrentDayName() time.Weekday {
	return time.Now().Weekday()
}
