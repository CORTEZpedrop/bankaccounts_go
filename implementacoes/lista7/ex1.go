package main

import (
	"fmt"
)

type DisjointSet struct {
	parent []int
}

// Função para criar uma nova coleção de conjuntos disjuntos.
func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parent: make([]int, n+1), // +1 para evitar índices baseados em zero
	}

	// Inicializa cada conjunto como um conjunto individual
	for i := 1; i <= n; i++ {
		ds.parent[i] = i
	}

	return ds
}

// Função para encontrar o representante (raiz) de um conjunto.
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		// Caminha recursivamente até encontrar o representante (raiz) do conjunto
		ds.parent[x] = ds.Find(ds.parent[x])
	}
	return ds.parent[x]
}

// Função para unir (fazer a união) de dois conjuntos.
func (ds *DisjointSet) Union(x, y int) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)

	if rootX != rootY {
		// Define o representante de um conjunto como o representante do outro conjunto
		ds.parent[rootX] = rootY
	}
}

// Função para verificar se dois elementos pertencem ao mesmo conjunto.
func (ds *DisjointSet) SameSet(x, y int) bool {
	return ds.Find(x) == ds.Find(y)
}

func main() {
	n := 10 // Substitua por seu valor de n

	// Crie uma nova coleção de conjuntos disjuntos com n elementos
	ds := NewDisjointSet(n)

	// Exemplo de uso das funções:
	ds.Union(1, 2)
	ds.Union(3, 4)
	ds.Union(5, 6)

	fmt.Println("Elementos 1 e 2 pertencem ao mesmo conjunto?", ds.SameSet(1, 2))
	fmt.Println("Elementos 1 e 3 pertencem ao mesmo conjunto?", ds.SameSet(1, 3))
	fmt.Println("Elementos 4 e 5 pertencem ao mesmo conjunto?", ds.SameSet(4, 5))
}
