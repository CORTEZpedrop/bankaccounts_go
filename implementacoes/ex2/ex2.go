package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type SortedLinkedList struct {
	Head *Node
}

func NewSortedLinkedList() *SortedLinkedList {
	return &SortedLinkedList{}
}

func (list *SortedLinkedList) Insert(value int) {
	newNode := &Node{Data: value}
	if list.Head == nil || value < list.Head.Data {
		newNode.Next = list.Head
		list.Head = newNode
		return
	}
	current := list.Head
	for current.Next != nil && current.Next.Data < value {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
}

func (list *SortedLinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func printRecursive(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Data)
	printRecursive(node.Next)
}

func (list *SortedLinkedList) PrintRecursive() {
	printRecursive(list.Head)
	fmt.Println()
}

func printReverseRecursive(node *Node) {
	if node == nil {
		return
	}
	printReverseRecursive(node.Next)
	fmt.Printf("%d ", node.Data)
}

func (list *SortedLinkedList) PrintReverseRecursive() {
	printReverseRecursive(list.Head)
	fmt.Println()
}

func (list *SortedLinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *SortedLinkedList) Search(value int) *Node {
	current := list.Head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (list *SortedLinkedList) Remove(value int) {
	if list.IsEmpty() {
		return
	}

	if list.Head.Data == value {
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	for current.Next != nil {
		if current.Next.Data == value {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func removeRecursive(prev *Node, current *Node, value int) {
	if current == nil {
		return
	}
	if current.Data == value {
		prev.Next = current.Next
		return
	}
	removeRecursive(current, current.Next, value)
}

func (list *SortedLinkedList) RemoveRecursive(value int) {
	if list.IsEmpty() {
		return
	}
	if list.Head.Data == value {
		list.Head = list.Head.Next
		return
	}
	removeRecursive(list.Head, list.Head.Next, value)
}

func (list *SortedLinkedList) Free() {
	list.Head = nil
}

func isEqual(list1 *SortedLinkedList, list2 *SortedLinkedList) bool {
	node1 := list1.Head
	node2 := list2.Head
	for node1 != nil && node2 != nil {
		if node1.Data != node2.Data {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return node1 == nil && node2 == nil
}

func main() {
	list := NewSortedLinkedList()

	list.Insert(3)
	list.Insert(1)
	list.Insert(2)

	fmt.Println("Valores na lista:")
	list.Print()

	fmt.Println("Valores na lista (recursão):")
	list.PrintRecursive()

	fmt.Println("Valores na lista em ordem reversa:")
	list.PrintReverseRecursive()

	fmt.Println("Buscar elemento 2:", list.Search(2))

	list.Remove(2)
	fmt.Println("Valores na lista após remover 2:")
	list.Print()

	fmt.Println("Buscar elemento 3:", list.Search(3))

	list.RemoveRecursive(3)
	fmt.Println("Valores na lista após remover 3 (recursão):")
	list.Print()

	list2 := NewSortedLinkedList()
	list2.Insert(1)
	list2.Insert(2)
	list2.Insert(3)

	fmt.Println("Lista 1 é igual à Lista 2:", isEqual(list, list2))

	list.Free()
	fmt.Println("Lista liberada")
	fmt.Println("Exercício 2")

}
