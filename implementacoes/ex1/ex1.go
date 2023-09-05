package main

import (
	"fmt"
)

// Definindo a estrutura do nó da lista encadeada
type Node struct {
	Data int
	Next *Node
}

// Definindo o tipo da lista encadeada
type LinkedList struct {
	Head *Node
}

// Função para criar uma lista vazia
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Função para inserir elemento no início da lista
func (list *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{Data: data, Next: list.Head}
	list.Head = newNode
}

// Função para imprimir os valores armazenados na lista
func (list *LinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

// Função para imprimir os valores usando recursão
func printRecursive(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Data)
	printRecursive(node.Next)
}

// Função para imprimir os valores em ordem reversa usando recursão
func printReverseRecursive(node *Node) {
	if node == nil {
		return
	}
	printReverseRecursive(node.Next)
	fmt.Printf("%d ", node.Data)
}

// Função para verificar se a lista está vazia
func (list *LinkedList) IsEmpty() bool {
	return list.Head == nil
}

// Função para buscar um elemento na lista
func (list *LinkedList) Search(value int) *Node {
	current := list.Head
	for current != nil {
		if current.Data == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// Função para remover um elemento da lista
func (list *LinkedList) Remove(value int) {
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

// Função para liberar a lista
func (list *LinkedList) Free() {
	list.Head = nil
}

func main() {
	list := NewLinkedList()

	list.InsertAtBeginning(3)
	list.InsertAtBeginning(2)
	list.InsertAtBeginning(1)

	fmt.Println("Valores na lista:")
	list.Print()

	fmt.Println("Valores na lista (recursão):")
	printRecursive(list.Head)
	fmt.Println()

	fmt.Println("Valores na lista em ordem reversa:")
	printReverseRecursive(list.Head)
	fmt.Println()

	fmt.Println("Buscar elemento 2:", list.Search(2))

	list.Remove(2)
	fmt.Println("Valores na lista após remover 2:")
	list.Print()

	list.Free()
	fmt.Println("Lista liberada")
	fmt.Println("exercicio 1")
}
