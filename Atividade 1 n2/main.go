package main

import "fmt"

type no struct {
	valor    int
	anterior *no
	proximo  *no
}

type lista struct {
	inicio *no
}

func (l *lista) inserirInicio(valor int) {
	novo := &no{valor: valor}

	if l.inicio == nil {
		l.inicio = novo
		return
	}

	novo.proximo = l.inicio
	l.inicio.anterior = novo
	l.inicio = novo
}

func (l *lista) inserirFim(valor int) {
	novo := &no{valor: valor}

	if l.inicio == nil {
		l.inicio = novo
		return
	}

	atual := l.inicio
	for atual.proximo != nil {
		atual = atual.proximo
	}

	atual.proximo = novo
	novo.anterior = atual
}

func (l *lista) removerHead() {
	if l.inicio == nil {
		fmt.Println("Lista já está vazia.")
		return
	}

	if l.inicio.proximo == nil {
		l.inicio = nil
		return
	}

	l.inicio = l.inicio.proximo
	l.inicio.anterior = nil
}

func (l *lista) removerTail() {
	if l.inicio == nil {
		fmt.Println("Lista já está vazia.")
		return
	}

	if l.inicio.proximo == nil {
		l.inicio = nil
		return
	}

	atual := l.inicio
	for atual.proximo != nil {
		atual = atual.proximo
	}

	atual.anterior.proximo = nil
}

func (l *lista) imprimir(mensagem string) {
	fmt.Printf("%s: ", mensagem)
	atual := l.inicio
	for atual != nil {
		fmt.Printf("[%d]", atual.valor)
		if atual.proximo != nil {
			fmt.Print(" <-> ")
		}
		atual = atual.proximo
	}
	fmt.Println()
}

func main() {
	l := lista{}

	l.inserirFim(10)
	l.inserirFim(20)
	l.inserirFim(50)
	l.inserirFim(60)
	l.inserirFim(80)

	l.imprimir("Lista Inicial")

	l.inserirInicio(5)
	l.imprimir("1. Inserido 5 no começo")

	l.inserirFim(5)
	l.imprimir("2. Inserido 5 no final (sem tail)")

	l.removerHead()
	l.imprimir("3. Removido o head")

	l.removerTail()
	l.imprimir("4. Removido o tail")
}
