package main

import (
	"fmt"
)

type Set struct {
	data map[string]bool
}

// Função para criar um novo conjunto.
func NewSet() *Set {
	return &Set{
		data: make(map[string]bool),
	}
}

// Função para inserir um elemento em um conjunto.
func (s *Set) Insert(element string) {
	s.data[element] = true
}

// Função para remover um elemento de um conjunto.
func (s *Set) Remove(element string) {
	delete(s.data, element)
}

// Função para fazer a união entre dois conjuntos.
func Union(setA, setB *Set) *Set {
	result := NewSet()

	for element := range setA.data {
		result.Insert(element)
	}
	for element := range setB.data {
		result.Insert(element)
	}

	return result
}

// Função para fazer a interseção entre dois conjuntos.
func Intersection(setA, setB *Set) *Set {
	result := NewSet()

	for element := range setA.data {
		if setB.Contains(element) {
			result.Insert(element)
		}
	}

	return result
}

// Função para fazer a diferença entre dois conjuntos.
func Difference(setA, setB *Set) *Set {
	result := NewSet()

	for element := range setA.data {
		if !setB.Contains(element) {
			result.Insert(element)
		}
	}

	return result
}

// Função para verificar se um conjunto A é subconjunto de B.
func (s *Set) IsSubsetOf(other *Set) bool {
	for element := range s.data {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// Função para verificar se dois conjuntos são iguais.
func (s *Set) IsEqual(other *Set) bool {
	if len(s.data) != len(other.data) {
		return false
	}

	for element := range s.data {
		if !other.Contains(element) {
			return false
		}
	}

	return true
}

// Função para gerar o complemento de um conjunto.
func (s *Set) Complement(universe *Set) *Set {
	result := NewSet()

	for element := range universe.data {
		if !s.Contains(element) {
			result.Insert(element)
		}
	}

	return result
}

// Função para verificar se um elemento pertence a um conjunto.
func (s *Set) Contains(element string) bool {
	_, exists := s.data[element]
	return exists
}

// Função para recuperar o número de elementos de um conjunto.
func (s *Set) Size() int {
	return len(s.data)
}

// Função para liberar a estrutura de dados (um determinado conjunto).
func (s *Set) Clear() {
	s.data = make(map[string]bool)
}

func main() {
	// Criar conjuntos
	setA := NewSet()
	setB := NewSet()

	// Inserir elementos
	setA.Insert("apple")
	setA.Insert("banana")
	setA.Insert("cherry")

	setB.Insert("banana")
	setB.Insert("date")
	setB.Insert("apple")

	// Imprimir conjuntos
	fmt.Println("Conjunto A:", setA.data)
	fmt.Println("Conjunto B:", setB.data)

	// União
	union := Union(setA, setB)
	fmt.Println("União (A U B):", union.data)

	// Interseção
	intersection := Intersection(setA, setB)
	fmt.Println("Interseção (A ∩ B):", intersection.data)

	// Diferença
	difference := Difference(setA, setB)
	fmt.Println("Diferença (A - B):", difference.data)

	// Verificar subconjunto
	fmt.Println("É subconjunto? A é subconjunto de B:", setA.IsSubsetOf(setB))

	// Verificar igualdade
	fmt.Println("São iguais? A é igual a B:", setA.IsEqual(setB))

	// Complemento
	universe := NewSet()
	universe.Insert("apple")
	universe.Insert("banana")
	universe.Insert("cherry")
	universe.Insert("date")
	fmt.Println("Conjunto Universo:", universe.data)
	complement := setA.Complement(universe)
	fmt.Println("Complemento de A:", complement.data)

	// Verificar pertencimento
	fmt.Println("Contém 'apple' em A:", setA.Contains("apple"))
	fmt.Println("Contém 'date' em A:", setA.Contains("date"))

	// Tamanho do conjunto
	fmt.Println("Tamanho de A:", setA.Size())

	// Limpar conjunto
	setA.Clear()
	fmt.Println("Conjunto A após limpeza:", setA.data)
}
