package main

import (
	"time"
)

func main() {
	nonNilErrorInterface()
	acceptInterfacesReturnStructs()
	deferWithErrors()
	deferMethods()
	safeRecover()

	time.Sleep(time.Second)
}
