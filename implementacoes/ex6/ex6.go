package main

import (
	"fmt"
)

// Conta é uma interface comum para todas as contas.
type Conta interface {
	GetNumero() int
	GetSaldo() float64
	EfetuarCredito(valor float64)
	EfetuarDebito(valor float64)
	String() string
}

type Node struct {
	conta Conta
	next  *Node
}

type ListaEncadeada struct {
	head *Node
}

func (lista *ListaEncadeada) InserirInicio(conta Conta) {
	novoNode := &Node{conta: conta, next: lista.head}
	lista.head = novoNode
}

func (lista *ListaEncadeada) Imprimir() {
	current := lista.head
	for current != nil {
		fmt.Println(current.conta)
		current = current.next
	}
}

func (lista *ListaEncadeada) VerificarVazia() bool {
	return lista.head == nil
}

func (lista *ListaEncadeada) BuscarConta(numero int) Conta {
	current := lista.head
	for current != nil {
		if current.conta.GetNumero() == numero {
			return current.conta
		}
		current = current.next
	}
	return nil
}

func (lista *ListaEncadeada) RemoverConta(numero int) {
	if lista.VerificarVazia() {
		return
	}
	if lista.head.conta.GetNumero() == numero {
		lista.head = lista.head.next
		return
	}
	prev := lista.head
	current := lista.head.next
	for current != nil {
		if current.conta.GetNumero() == numero {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}

func (lista *ListaEncadeada) Liberar() {
	lista.head = nil
}

type ContaBase struct {
	numero int
	saldo  float64
}

func (c *ContaBase) GetNumero() int {
	return c.numero
}

func (c *ContaBase) GetSaldo() float64 {
	return c.saldo
}

func (c *ContaBase) EfetuarCredito(valor float64) {
	c.saldo += valor
}

func (c *ContaBase) EfetuarDebito(valor float64) {
	if c.saldo >= valor {
		c.saldo -= valor
	}
}

func (c *ContaBase) String() string {
	return fmt.Sprintf("Conta %d - Saldo: %.2f", c.numero, c.saldo)
}

type ContaPoupanca struct {
	ContaBase
}

func (c *ContaPoupanca) RenderJuros(taxa float64) {
	c.saldo += c.saldo * taxa
}

type ContaFidelidade struct {
	ContaBase
	bonus float64
}

func (c *ContaFidelidade) EfetuarCredito(valor float64) {
	c.ContaBase.EfetuarCredito(valor)
	c.bonus += valor * 0.01
}

func (c *ContaFidelidade) RenderBonus() {
	c.EfetuarCredito(c.bonus)
	c.bonus = 0
}

func main() {
	listaContas := ListaEncadeada{}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Inserir Conta Bancária")
		fmt.Println("2. Inserir Conta Poupança")
		fmt.Println("3. Inserir Conta Fidelidade")
		fmt.Println("4. Realizar crédito")
		fmt.Println("5. Realizar débito")
		fmt.Println("6. Consultar saldo")
		fmt.Println("7. Consultar bonus")
		fmt.Println("8. Realizar transferência")
		fmt.Println("9. Render juros de poupanca")
		fmt.Println("10. Render bonus de uma conta fidelidade ")
		fmt.Println("11. Remover conta")
		fmt.Println("12. Imprimir todas contas e saldos")

		var escolha int
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			// Inserir Conta Bancária
			// Solicitar número e saldo, criar a conta e inserir na lista
			fmt.Println("Inserir Conta Bancária")
			var numeroConta int
			var saldoConta float64

			fmt.Print("Número da Conta: ")
			fmt.Scan(&numeroConta)

			fmt.Print("Saldo Inicial: ")
			fmt.Scan(&saldoConta)

			novaConta := &ContaBase{
				numero: numeroConta,
				saldo:  saldoConta,
			}

			listaContas.InserirInicio(novaConta)

			fmt.Println("Conta Bancária inserida com sucesso!")
		case 2:
			// Inserir Conta Poupança
			// Solicitar número e saldo, criar a conta e inserir na lista
			fmt.Println("Inserir Conta Poupança")
			var numeroConta int
			var saldoConta float64

			fmt.Print("Número da Conta: ")
			fmt.Scan(&numeroConta)

			fmt.Print("Saldo Inicial: ")
			fmt.Scan(&saldoConta)

			novaContaPoupanca := &ContaPoupanca{
				ContaBase: ContaBase{
					numero: numeroConta,
					saldo:  saldoConta,
				},
			}

			listaContas.InserirInicio(novaContaPoupanca)

			fmt.Println("Conta Poupança inserida com sucesso!")
		case 3:
			// Inserir Conta Fidelidade
			// Solicitar número, saldo e bônus, criar a conta e inserir na lista
			fmt.Println("Inserir Conta Fidelidade")
			var numeroConta int
			var saldoConta float64
			var bonusConta float64

			fmt.Print("Número da Conta: ")
			fmt.Scan(&numeroConta)

			fmt.Print("Saldo Inicial: ")
			fmt.Scan(&saldoConta)

			fmt.Print("Bônus Inicial: ")
			fmt.Scan(&bonusConta)

			novaContaFidelidade := &ContaFidelidade{
				ContaBase: ContaBase{
					numero: numeroConta,
					saldo:  saldoConta,
				},
				bonus: bonusConta,
			}

			listaContas.InserirInicio(novaContaFidelidade)

			fmt.Println("Conta Fidelidade inserida com sucesso!")
		case 4:
			// Realizar Crédito em uma Conta
			// Solicitar número da conta e valor a ser creditado, realizar a operação
			fmt.Println("Realizar Crédito em uma Conta")
			var numeroConta int
			var valorCredito float64

			fmt.Print("Número da Conta: ")
			fmt.Scan(&numeroConta)

			fmt.Print("Valor a ser creditado: ")
			fmt.Scan(&valorCredito)

			conta := listaContas.BuscarConta(numeroConta)
			if conta != nil {
				conta.EfetuarCredito(valorCredito)
				fmt.Println("Crédito realizado com sucesso.")
			} else {
				fmt.Println("Conta não encontrada.")
			}
		case 5:
			// Realizar Débito em uma Conta
			// Solicitar número da conta e valor a ser debitado, realizar a operação
			fmt.Println("Realizar Débito em uma Conta")
			var numeroConta int
			var valorDebito float64

			fmt.Print("Número da Conta: ")
			fmt.Scan(&numeroConta)

			fmt.Print("Valor a ser debitado: ")
			fmt.Scan(&valorDebito)

			conta := listaContas.BuscarConta(numeroConta)
			if conta != nil {
				conta.EfetuarDebito(valorDebito)
				fmt.Println("Débito realizado com sucesso.")
			} else {
				fmt.Println("Conta não encontrada.")
			}
		case 6:
			// Consultar Saldo de uma Conta
			// Solicitar número da conta, buscar e imprimir o saldo
			fmt.Println("Consultar Saldo de uma Conta")
			var numeroConta int

			fmt.Print("Número da Conta: ")
			fmt.Scan(&numeroConta)

			conta := listaContas.BuscarConta(numeroConta)
			if conta != nil {
				saldo := conta.GetSaldo()
				fmt.Printf("Saldo da Conta: %.2f\n", saldo)
			} else {
				fmt.Println("Conta não encontrada.")
			}
		case 7:
			// Consultar Bônus de uma Conta Fidelidade
			// Solicitar número da conta fidelidade, buscar e imprimir o bônus
			fmt.Println("Consultar Bônus de uma Conta Fidelidade")
			var numeroConta int

			fmt.Print("Número da Conta Fidelidade: ")
			fmt.Scan(&numeroConta)

			conta := listaContas.BuscarConta(numeroConta)
			if contaFidelidade, ok := conta.(*ContaFidelidade); ok {
				bonus := contaFidelidade.bonus
				fmt.Printf("Bônus da Conta Fidelidade: %.2f\n", bonus)
			} else {
				fmt.Println("Conta Fidelidade não encontrada.")
			}
		case 8:
			// Realizar Transferência entre Contas
			// Solicitar números das contas de origem e destino, e valor a ser transferido
			fmt.Println("Realizar Transferência entre Duas Contas")
			var numeroContaOrigem, numeroContaDestino int
			var valorTransferencia float64

			fmt.Print("Número da Conta de Origem: ")
			fmt.Scan(&numeroContaOrigem)

			fmt.Print("Número da Conta de Destino: ")
			fmt.Scan(&numeroContaDestino)

			fmt.Print("Valor a ser Transferido: ")
			fmt.Scan(&valorTransferencia)

			contaOrigem := listaContas.BuscarConta(numeroContaOrigem)
			contaDestino := listaContas.BuscarConta(numeroContaDestino)

			if contaOrigem != nil && contaDestino != nil {
				if saldoOrigem := contaOrigem.GetSaldo(); saldoOrigem >= valorTransferencia {
					contaOrigem.EfetuarDebito(valorTransferencia)
					contaDestino.EfetuarCredito(valorTransferencia)
					fmt.Println("Transferência realizada com sucesso.")
				} else {
					fmt.Println("Saldo insuficiente na Conta de Origem.")
				}
			} else {
				fmt.Println("Conta de Origem e/ou Conta de Destino não encontradas.")
			}
		case 9:
			// Render Juros de uma Conta Poupança
			// Solicitar número da conta poupança e taxa de juros, realizar a operação
			fmt.Println("Render Juros de uma Conta Poupança")
			var numeroContaPoupanca int
			var taxaJuros float64

			fmt.Print("Número da Conta Poupança: ")
			fmt.Scan(&numeroContaPoupanca)

			fmt.Print("Taxa de Juros (%): ")
			fmt.Scan(&taxaJuros)

			contaPoupanca := listaContas.BuscarConta(numeroContaPoupanca)

			if contaPoupanca != nil {
				if contaPoupancaPoupanca, ok := contaPoupanca.(*ContaPoupanca); ok {
					contaPoupancaPoupanca.RenderJuros(taxaJuros / 100.0)
					fmt.Println("Juros renderizados com sucesso.")
				} else {
					fmt.Println("A conta especificada não é uma Conta Poupança.")
				}
			} else {
				fmt.Println("Conta Poupança não encontrada.")
			}
		case 10:
			// Render Bônus de uma Conta Fidelidade
			// Solicitar número da conta fidelidade, realizar a operação
			fmt.Println("Render Bônus de uma Conta Fidelidade")
			var numeroContaFidelidade int

			fmt.Print("Número da Conta Fidelidade: ")
			fmt.Scan(&numeroContaFidelidade)

			contaFidelidade := listaContas.BuscarConta(numeroContaFidelidade)

			if contaFidelidade != nil {
				if contaFidelidadeFidelidade, ok := contaFidelidade.(*ContaFidelidade); ok {
					contaFidelidadeFidelidade.RenderBonus()
					fmt.Println("Bônus renderizado com sucesso.")
				} else {
					fmt.Println("A conta especificada não é uma Conta Fidelidade.")
				}
			} else {
				fmt.Println("Conta Fidelidade não encontrada.")
			}
		case 11:
			// Remover uma Conta
			// Solicitar número da conta, remover da lista
			fmt.Println("Remover uma Conta")
			var numeroContaRemover int

			fmt.Print("Número da Conta a ser Removida: ")
			fmt.Scan(&numeroContaRemover)

			listaContas.RemoverConta(numeroContaRemover)

			fmt.Println("Conta removida com sucesso, se existia.")
		case 12:
			// Imprimir número e saldo de todas as contas cadastradas
			listaContas.Imprimir()
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
