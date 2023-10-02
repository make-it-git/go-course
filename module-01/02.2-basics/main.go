package main

import (
	"fmt"

	"02.2-basics/example"
	alias "02.2-basics/example2"
)

func main() {
	fmt.Println(example.MyExportedValue)
	fmt.Println(alias.MyOtherValue)
}
