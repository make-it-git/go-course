package main

import (
	"fmt"
)

func main() {
	noGenerics()

	// Ограничения снизу (по типам)
	fmt.Println(">>> type constrains")
	typeConstraints()

	// Ограничения сверху (по интерфейсам)
	fmt.Println(">>> interface constraints")
	interfaceConstraints()

	// Встроенные constrains
	fmt.Println(">>> builtin constraints")
	builtinConstraints()

	fmt.Println(">>> custom constraints")
	customConstraints()

	fmt.Println(">>> examples")
	fmt.Println("> linked list")
	linkedListExample()
	fmt.Println("> stack")
	stackExample()
	fmt.Println("> functional")
	functionalExamples()
	fmt.Println("> decorator")
	decoratorExample()

	fmt.Println(">>> func range types")
	funcRangeExample()
}
