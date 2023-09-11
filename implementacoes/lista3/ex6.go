package main

import (
	"fmt"
	"os"
)

type Node struct {
	key  int
	next *Node
}

type HashTable struct {
	size    int
	buckets []*Node
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([]*Node, size),
	}
}

func (ht *HashTable) hash(key int) int {
	return key % ht.size
}

func (ht *HashTable) Insert(key int) {
	index := ht.hash(key)
	newNode := &Node{key: key, next: nil}

	if ht.buckets[index] == nil {
		ht.buckets[index] = newNode
	} else {
		current := ht.buckets[index]
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ht *HashTable) Search(key int) bool {
	index := ht.hash(key)
	current := ht.buckets[index]

	for current != nil {
		if current.key == key {
			return true
		}
		current = current.next
	}

	return false
}

func (ht *HashTable) Remove(key int) {
	index := ht.hash(key)
	current := ht.buckets[index]
	var prev *Node = nil

	for current != nil {
		if current.key == key {
			if prev == nil {
				ht.buckets[index] = current.next
			} else {
				prev.next = current.next
			}
			return
		}
		prev = current
		current = current.next
	}
}

func (ht *HashTable) Print() {
	for i := 0; i < ht.size; i++ {
		fmt.Printf("Bucket %d: ", i)
		current := ht.buckets[i]
		for current != nil {
			fmt.Printf("%d -> ", current.key)
			current = current.next
		}
		fmt.Println("nil")
	}
}

func (ht *HashTable) Free() {
	ht.buckets = make([]*Node, ht.size)
}

func main() {
	fmt.Print("Lista de exercicios numero 2 \n")
	var n int
	fmt.Print("Informe o valor de n: ")
	fmt.Scan(&n)
	if n < 2 {
		fmt.Println("Escolha um valor maior que 1.")
		os.Exit(0)
	}

	table := NewHashTable(n / 2)

	for {
		fmt.Println("\nEscolha uma operação:")
		fmt.Println("1. Inserir elemento")
		fmt.Println("2. Buscar elemento")
		fmt.Println("3. Remover elemento")
		fmt.Println("4. Imprimir tabela")
		fmt.Println("5. Sair")
		fmt.Print("Opção: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Informe o número a ser inserido: ")
			var key int
			fmt.Scan(&key)
			table.Insert(key)
			fmt.Printf("Chave %d inserida na tabela.\n", key)

		case 2:
			fmt.Print("Informe o número a ser buscado: ")
			var key int
			fmt.Scan(&key)
			if table.Search(key) {
				fmt.Printf("Chave %d encontrada na tabela.\n", key)
			} else {
				fmt.Printf("Chave %d não encontrada na tabela.\n", key)
			}

		case 3:
			fmt.Print("Informe o número a ser removido: ")
			var key int
			fmt.Scan(&key)
			table.Remove(key)
			fmt.Printf("Chave %d removida da tabela.\n", key)

		case 4:
			fmt.Println("Tabela de Dispersão:")
			table.Print()

		case 5:
			fmt.Println("Saindo do programa.")
			os.Exit(0)

		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
