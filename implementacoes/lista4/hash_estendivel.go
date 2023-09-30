package main

import (
	"fmt"
	"math"
)

// Definindo uma estrutura para o bucket
type Bucket struct {
	elements []int
}

// Definindo a estrutura para a Hash Estendível
type ExtensibleHash struct {
	n          int       // valor máximo da chave
	bucketSize int       // tamanho máximo de um bucket
	buckets    []*Bucket // slice de buckets
	depth      int       // profundidade global
}

// Função para criar a estrutura de dados (hash estendível)
func NewExtensibleHash(n, bucketSize int) *ExtensibleHash {
	depth := int(math.Log2(float64(n)))
	buckets := make([]*Bucket, 1<<depth)
	for i := range buckets {
		buckets[i] = &Bucket{}
	}
	return &ExtensibleHash{
		n:          n,
		bucketSize: bucketSize,
		buckets:    buckets,
		depth:      depth,
	}
}

// Função para inserir um elemento na estrutura
func (eh *ExtensibleHash) Insert(key int) {
	index := key % (1 << eh.depth)
	bucket := eh.buckets[index]
	if len(bucket.elements) < eh.bucketSize {
		bucket.elements = append(bucket.elements, key)
	} else {
		// Realizar divisão de buckets
		if eh.depth == 0 {
			eh.depth++
			eh.buckets = append(eh.buckets, &Bucket{}, &Bucket{})
			for _, value := range bucket.elements {
				idx := value % (1 << eh.depth)
				eh.buckets[idx].elements = append(eh.buckets[idx].elements, value)
			}
			bucket.elements = nil
			eh.Insert(key) // Recursivamente tenta inserir novamente
		}
	}
}

// Função para recuperar/buscar um elemento na estrutura
func (eh *ExtensibleHash) Search(key int) bool {
	index := key % (1 << eh.depth)
	bucket := eh.buckets[index]
	for _, element := range bucket.elements {
		if element == key {
			return true
		}
	}
	return false
}

// Função para remover um elemento da estrutura
func (eh *ExtensibleHash) Remove(key int) {
	index := key % (1 << eh.depth)
	bucket := eh.buckets[index]
	for i, element := range bucket.elements {
		if element == key {
			bucket.elements = append(bucket.elements[:i], bucket.elements[i+1:]...)
			break
		}
	}
}

// Função para liberar a estrutura de dados (hash estendível)
func (eh *ExtensibleHash) Destroy() {
	eh.buckets = nil
}

func main() {
	n := 100        // Valor máximo da chave
	bucketSize := 5 // Tamanho máximo de um bucket

	// Criar a estrutura de dados
	eh := NewExtensibleHash(n, bucketSize)

	// Inserir elementos
	for i := 1; i <= n; i++ {
		eh.Insert(i)
	}

	// Buscar um elemento
	keyToSearch := 42
	if eh.Search(keyToSearch) {
		fmt.Printf("Elemento %d encontrado.\n", keyToSearch)
	} else {
		fmt.Printf("Elemento %d não encontrado.\n", keyToSearch)
	}

	// Remover um elemento
	keyToRemove := 42
	eh.Remove(keyToRemove)
	fmt.Printf("Elemento %d removido.\n", keyToRemove)

	// Liberar a estrutura de dados
	eh.Destroy()
}
