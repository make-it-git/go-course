package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (ll *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{data: value}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *LinkedList[T]) Print() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func linkedListExample() {
	intList := LinkedList[int]{}
	intList.Add(10)
	intList.Add(20)
	intList.Add(30)
	intList.Print()

	stringList := LinkedList[string]{}
	stringList.Add("apple")
	stringList.Add("banana")
	stringList.Add("cherry")
	stringList.Print()
}
