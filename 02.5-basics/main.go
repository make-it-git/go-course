package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name     string
	Birthday time.Time
}

// func (variable T) methodName(params) (returnTypes) {}

func (c Person) doNothing() {

}

func (c Person) CanDriveCar() bool {
	return c.Birthday.AddDate(18, 0, 0).Before(time.Now())
}

func (c Person) UpdateName(name string) {
	c.Name = name
}

func (c *Person) UpdateNameByPointer(name string) {
	c.Name = name
}

type MyInt int

func (i MyInt) isGreater(value int) bool {
	return i > MyInt(value)
}

func main() {
	const layout = "2006-Jan-02"
	tm, _ := time.Parse(layout, "2000-Sep-03")
	p := Person{
		Name:     "John",
		Birthday: tm,
	}
	fmt.Println(p) // {John 2000-02-01 00:00:00 +0000 UTC}

	p.UpdateName("new name")
	fmt.Println(p) // unchanged

	p.UpdateNameByPointer("new name") // (&p).UpdateNameByPointer(...)
	fmt.Println(p)                    // {new name 2000-02-01 00:00:00 +0000 UTC}
	fmt.Println(p.CanDriveCar())      // true

	i := MyInt(10)
	fmt.Println(i.isGreater(10)) // false
}
