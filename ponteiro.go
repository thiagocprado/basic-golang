package main

import "fmt"

func main() {
	somaSemPonteiro()
	somaComPonteiro()
}

// Quando passamos um valor para uma função sem usar um ponteiro,
// você está passando uma cópia desse valor. Isso significa que qualquer
// modificação feita dentro da função não afetará o valor original fora dela.

// Neste exemplo, quando chamamos incrementa(x), x é passado como um valor para a função
// incrementa. Dentro da função, num é uma cópia do valor de x. Qualquer modificação feita
// em num dentro da função não afeta o valor original de x. Portanto, quando imprimimos x
// novamente no somaSemPonteiro(), seu valor permanece o mesmo, ou seja, 10.
func somaSemPonteiro() {
	x := 10
	fmt.Println("Antes da função:", x)
	incrementa(x)
	fmt.Println("Depois da função (sem ponteiro):", x)
	fmt.Println("")
}

func incrementa(num int) {
	num++
}

// Aqui, ao chamar incrementaComPonteiro(&x), estamos passando o endereço de memória onde x
// está armazenado (&x). Dentro da função incrementaComPonteiro, num é um ponteiro para int,
// e ao usar *num++, estamos incrementando o valor naquela posição de memória, que é o mesmo
// que x no somaComPonteiro().

// Portanto, quando imprimimos x após chamar incrementaComPonteiro, vemos que seu valor foi
// alterado para 11
func somaComPonteiro() {
	x := 10
	fmt.Println("Antes da função:", x)
	incrementaComPonteiro(&x)
	fmt.Println("Depois da função (com ponteiro):", x)
	fmt.Println("")
}

func incrementaComPonteiro(num *int) {
	*num++
}

// Conclusão

// Ponteiro em Go é como um bilhete que aponta para onde um valor está guardado
// na memória do computador. Ele permite que você passe e modifique valores diretamente
// sem ter que copiar muitos dados, o que pode fazer seu programa mais rápido e eficiente.

// Vantagens

// Controle de Memória:
// - Eficiência: Usar ponteiros permite evitar cópias desnecessárias de grandes estruturas de dados, o que pode melhorar a eficiência do programa.
// - Mutabilidade: Facilita a alteração de dados em uma função chamada, permitindo que você passe uma referência a uma estrutura de dados e modifique seu conteúdo diretamente.

// Modelagem de Estruturas Complexas:
// -Estruturas de Dados: Ponteiros são essenciais para a criação de estruturas de dados como listas encadeadas, árvores, e gráficos, onde os nós precisam referenciar outros nós.

// Interoperabilidade:
// - C APIs: Trabalhar com ponteiros é fundamental para interoperar com bibliotecas escritas em C ou outras linguagens que utilizam ponteiros.

// Semântica Clara de Propriedade:
// -Transparência: O uso explícito de ponteiros pode deixar claro que uma função pode modificar os dados passados, o que pode ser mais intuitivo para quem está lendo o código.

// Desvantagens

// Complexidade Adicional:
// - Gerenciamento de Memória: Trabalhar com ponteiros requer um entendimento cuidadoso de como a memória é gerenciada, o que pode aumentar a complexidade do código.
// - Erros Difíceis de Diagnosticar: Erros como acessos inválidos, vazamentos de memória, e uso de memória após liberação (dangling pointers) podem ser difíceis de diagnosticar e corrigir.

// Segurança:
// - Segurança: Ponteiros podem levar a problemas de segurança, como vulnerabilidades de buffer overflow, que são menos prováveis em linguagens que abstraem o gerenciamento de memória.

// Legibilidade:
// - Código Mais Difícil de Ler: O uso intensivo de ponteiros pode tornar o código mais difícil de ler e entender, especialmente para desenvolvedores menos experientes.

// Desempenho Não Garantido:
// - Overhead de Indireção: Embora ponteiros possam melhorar a eficiência, o uso excessivo pode introduzir overhead de indireção, o que pode, em alguns casos, prejudicar o desempenho.
