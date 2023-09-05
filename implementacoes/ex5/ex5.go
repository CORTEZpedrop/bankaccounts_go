package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type CircularDoublyLinkedList struct {
	Head *Node
}

func NewCircularDoublyLinkedList() *CircularDoublyLinkedList {
	return &CircularDoublyLinkedList{}
}

func (list *CircularDoublyLinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *CircularDoublyLinkedList) Insert(value int) {
	newNode := &Node{Data: value}
	if list.IsEmpty() {
		newNode.Next = newNode
		newNode.Prev = newNode
		list.Head = newNode
		return
	}
	tail := list.Head.Prev
	tail.Next = newNode
	newNode.Prev = tail
	newNode.Next = list.Head
	list.Head.Prev = newNode
}

func (list *CircularDoublyLinkedList) Print() {
	if list.IsEmpty() {
		return
	}
	current := list.Head
	for {
		fmt.Printf("%d ", current.Data)
		current = current.Next
		if current == list.Head {
			break
		}
	}
	fmt.Println()
}

func printRecursive(node *Node, head *Node) {
	if node == head {
		return
	}
	fmt.Printf("%d ", node.Data)
	printRecursive(node.Next, head)
}

func (list *CircularDoublyLinkedList) PrintRecursive() {
	if list.IsEmpty() {
		return
	}
	printRecursive(list.Head, list.Head)
	fmt.Println()
}

func (list *CircularDoublyLinkedList) Search(value int) *Node {
	if list.IsEmpty() {
		return nil
	}
	current := list.Head
	for {
		if current.Data == value {
			return current
		}
		current = current.Next
		if current == list.Head {
			break
		}
	}
	return nil
}

func (list *CircularDoublyLinkedList) Remove(value int) {
	if list.IsEmpty() {
		return
	}
	if list.Head.Data == value {
		if list.Head.Next == list.Head {
			list.Head = nil
			return
		}
		tail := list.Head.Prev
		list.Head = list.Head.Next
		list.Head.Prev = tail
		tail.Next = list.Head
		return
	}
	current := list.Head.Next
	for current != list.Head {
		if current.Data == value {
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
			return
		}
		current = current.Next
	}
}

func removeRecursive(prev *Node, current *Node, value int, head *Node) {
	if current == head {
		return
	}
	if current.Data == value {
		prev.Next = current.Next
		current.Next.Prev = prev
		return
	}
	removeRecursive(current, current.Next, value, head)
}

func (list *CircularDoublyLinkedList) RemoveRecursive(value int) {
	if list.IsEmpty() {
		return
	}
	if list.Head.Data == value {
		if list.Head.Next == list.Head {
			list.Head = nil
			return
		}
		tail := list.Head.Prev
		list.Head = list.Head.Next
		list.Head.Prev = tail
		tail.Next = list.Head
		return
	}
	removeRecursive(list.Head, list.Head.Next, value, list.Head)
}

func (list *CircularDoublyLinkedList) Free() {
	list.Head = nil
}

func main() {
	list := NewCircularDoublyLinkedList()

	list.Insert(3)
	list.Insert(1)
	list.Insert(2)

	fmt.Println("Valores na lista:")
	list.Print()

	fmt.Println("Valores na lista (recursão):")
	list.PrintRecursive()

	fmt.Println("Buscar elemento 2:", list.Search(2))

	list.Remove(2)
	fmt.Println("Valores na lista após remover 2:")
	list.Print()

	fmt.Println("Buscar elemento 3:", list.Search(3))

	list.RemoveRecursive(3)
	fmt.Println("Valores na lista após remover 3 (recursão):")
	list.Print()

	list.Free()
	fmt.Println("Lista liberada")
	fmt.Println("exercicio 5")
}
