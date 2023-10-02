package main

import (
	"fmt"
)

func main() {
	ranges()
	shadowingVariables()
	typeEmbedding()
	valueReceivers()
	nilProblem()
	stringsImmutable()
	staleSlices()
}

// ranges
func ranges() {
	valuesStr := []string{"a", "b", "c"}
	for index, value := range valuesStr {
		fmt.Println(index, value) // 0 a, 1 b, etc.
	}

	valuesInt := []int{4, 8, 15, 16, 23, 42}
	for value := range valuesInt {
		fmt.Println(value) // 0 1 2, etc.
	}
}

// shadowing variables
func shadowingVariables() {
	node := Node{
		value:  "original",
		weight: 1,
	}

	t1 := func(n Node) (Node, error) {
		return Node{
			value:  n.value + " t1",
			weight: n.weight + 1,
		}, nil
	}

	t2 := func(n Node) (Node, error) {
		return Node{
			value:  n.value + " t2",
			weight: n.weight + 2,
		}, nil
	}

	transformed, err := applyTransformations(node, []Transformation{t1, t2})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", transformed) // [{original t1 2} {original t2 3}]
	}

	transformed2, err := applyTransformationsValid(node, []Transformation{t1, t2})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", transformed2) // [{original t1 2} {original t1 t2 4}]
	}
}

type Node struct {
	value  string
	weight int
}

type Transformation func(Node) (Node, error)

func applyTransformations(n Node, ts []Transformation) ([]Node, error) {
	var steps []Node
	for _, t := range ts {
		n, err := t(n) // shadowed variable
		if err != nil {
			return nil, err
		}
		steps = append(steps, n)
	}
	return steps, nil
}

func applyTransformationsValid(n Node, ts []Transformation) ([]Node, error) {
	var steps []Node
	for _, t := range ts {
		newNode, err := t(n)
		n = newNode
		if err != nil {
			return nil, err
		}
		steps = append(steps, n)
	}
	return steps, nil
}

// type embedding
func typeEmbedding() {
	v1 := Wrapper{Node{"123", 0}}
	fmt.Println(v1.Name(), v1.Node.Name()) // 123 123
	fmt.Println(v1.value, v1.Node.value)   // 123 123
	fmt.Println(v1.Name())                 // wrapped value: 123
	fmt.Println(v1.Node.Name())            // 123
}

func (n Node) Name() string {
	return n.value
}

type Wrapper struct {
	Node
}

func (w Wrapper) Name() string {
	return "wrapped value: " + w.Node.value
}

// value receivers
func valueReceivers() {
	// func (n Node) Name() string {...}   // Value Receiver
	// func (n *Node) Name() string {...}  // Pointer Receiver

	n := Node{value: "my value"}
	n.SetValue("new value")
	fmt.Println(n.value) // my value
}

func (n Node) SetValue(newValue string) {
	n.value = newValue
}

// nil
func nilProblem() {
	var slice []int
	fmt.Println(len(slice))       // 0
	slice = append(slice, 100500) // ok
	fmt.Println(len(slice))       // 1

	var m map[string]int
	fmt.Println(len(m)) // 0
	// m["test"] = 1       // panic: assignment to entry in nil map

	mOk := make(map[string]int)
	fmt.Println(len(mOk)) // 0
	mOk["test"] = 1
	fmt.Println(len(mOk)) // 1
}

// string immutable
func stringsImmutable() {
	s := "abc"
	// s[1] = 'B' // cannot assign to s[1] (value of type byte)

	sBytes := []byte(s)
	sBytes[1] = 'B'

	fmt.Println(string(sBytes)) // aBc

	sRunes := []rune(s)
	sRunes[1] = 'B'
	fmt.Println(string(sRunes)) // aBc

	s2 := "🐧🐧"
	r2 := []rune(s2)
	fmt.Println(string(r2)) // 🐧🐧
	r2[0] = 'A'
	fmt.Println(string(r2)) // A🐧
}

func staleSlices() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) // 3 3 [1 2 3]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) // 2 2 [2 3]

	s2[0] = 100500

	fmt.Println(s1) // [1 100500 3]
	fmt.Println(s2) // [100500 3]

	s2 = append(s2, 4)
	s2[0] = 10

	fmt.Println(s1) // [1 100500 3]
	fmt.Println(s2) // [10 3 4]
}
