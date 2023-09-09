package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	type Person struct {
		Name, Surname string
		Birthday      time.Time
	}

	date := time.Date(1970, 1, 1, 1, 1, 1, 1, time.Local)

	var p1 Person
	fmt.Println(p1) // {  0001-01-01 00:00:00 +0000 UTC}

	var p2 = Person{
		Name:     "John",
		Surname:  "Wick",
		Birthday: date,
	}
	fmt.Println(p2)

	var p3 = Person{Name: "Somebody"}
	fmt.Println(p3) // {Somebody  0001-01-01 00:00:00 +0000 UTC}

	var p4 = Person{"John", "Wick", date}
	fmt.Println(p4.Name) // John

	ptr := &p4
	fmt.Println(ptr.Name, (*ptr).Name) // John John

	fmt.Println(p2 == *ptr, p2 == p4)

	type SecretPerson struct {
		Name    string
		surname string
	}

	type Worker struct {
		Person
		Salary int
	}
	w := Worker{
		Person: Person{
			Name:     "John",
			Surname:  "Wick",
			Birthday: date,
		},
		Salary: 100500,
	}
	fmt.Println(w)                     // {{John Wick 1970-01-01 01:01:01.000000001 +0400 +04} 100500
	fmt.Println(w.Name, w.Person.Name) // John John

	type PersonWithTags struct {
		Name string `json:"name"`
	}

	type Point struct {
		X, Y float64
	}

	point1 := Point{1, 2}
	point2 := point1 // Copy of p1 is assigned to p2

	point2.X = 2

	fmt.Println(point1) // {1 2}
	fmt.Println(point2) // {2 2}

	type empty struct{}
	fmt.Println(unsafe.Sizeof(empty{})) // 0
}
