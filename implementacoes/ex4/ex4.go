package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type CircularLinkedList struct {
	Head *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (list *CircularLinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *CircularLinkedList) Insert(value int) {
	newNode := &Node{Data: value}
	if list.IsEmpty() {
		newNode.Next = newNode
		list.Head = newNode
		return
	}
	current := list.Head
	for current.Next != list.Head {
		current = current.Next
	}
	current.Next = newNode
	newNode.Next = list.Head
}

func (list *CircularLinkedList) Print() {
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

func (list *CircularLinkedList) PrintRecursive() {
	if list.IsEmpty() {
		return
	}
	printRecursive(list.Head, list.Head)
	fmt.Println()
}

func (list *CircularLinkedList) Search(value int) *Node {
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

func (list *CircularLinkedList) Remove(value int) {
	if list.IsEmpty() {
		return
	}
	if list.Head.Data == value {
		if list.Head.Next == list.Head {
			list.Head = nil
			return
		}
		current := list.Head
		for current.Next != list.Head {
			current = current.Next
		}
		list.Head = list.Head.Next
		current.Next = list.Head
		return
	}
	prev := list.Head
	current := list.Head.Next
	for current != list.Head {
		if current.Data == value {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}
}

func removeRecursive(prev *Node, current *Node, value int, head *Node) {
	if current == head {
		return
	}
	if current.Data == value {
		prev.Next = current.Next
		return
	}
	removeRecursive(current, current.Next, value, head)
}

func (list *CircularLinkedList) RemoveRecursive(value int) {
	if list.IsEmpty() {
		return
	}
	if list.Head.Data == value {
		if list.Head.Next == list.Head {
			list.Head = nil
			return
		}
		current := list.Head
		for current.Next != list.Head {
			current = current.Next
		}
		list.Head = list.Head.Next
		current.Next = list.Head
		return
	}
	removeRecursive(list.Head, list.Head.Next, value, list.Head)
}

func (list *CircularLinkedList) Free() {
	list.Head = nil
}

func main() {
	list := NewCircularLinkedList()

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
	fmt.Println("exercicio 4")
}
