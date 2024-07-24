package main

import (
	"fmt"
	"sync"
)

// Função que calcula a soma de uma porção do slice de números
func sum(numbers []int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done() // Indica que a goroutine terminou ao final da função
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	resultChan <- sum // Envia o resultado da soma para o canal
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 2) // Canal para receber resultados parciais, buffer de tamanho 2
	var wg sync.WaitGroup

	mid := len(numbers) / 2 // Divide o slice em duas partes aproximadamente iguais

	// Adiciona 2 ao WaitGroup para esperar 2 goroutines
	wg.Add(2)

	// Inicia duas goroutines para calcular a soma de cada metade do slice
	go sum(numbers[:mid], &wg, resultChan)
	go sum(numbers[mid:], &wg, resultChan)

	// Espera até que todas as goroutines terminem
	wg.Wait()

	// Fecha o canal após todas as goroutines terem terminado
	close(resultChan)

	totalSum := 0
	// Soma os resultados parciais recebidos do canal
	for partialSum := range resultChan {
		totalSum += partialSum
	}

	fmt.Printf("A soma total é: %d\n", totalSum)
}
