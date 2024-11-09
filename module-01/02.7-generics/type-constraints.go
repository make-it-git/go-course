package main

type additive interface {
	int | ~uint64
}

func add[T additive](a T, b T) T {
	return a + b
}

// Constraint literal
func add2[T int | ~uint64](a T, b T) T {
	return a + b
}

type iAmUnit64Alias uint64

func typeConstraints() {
	_ = add(1, 2)
	_ = add2(3, 4)

	_ = add(1, iAmUnit64Alias(2))
	_ = add2(3, iAmUnit64Alias(4))
}
