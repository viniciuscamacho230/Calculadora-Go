package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "regexp"

var reader = bufio.NewReader(os.Stdin)
var results = []string{}

func main() {
	for {
		fmt.Println("Seja bem-vindo a maravilhosa calculadora Golang")
		oneValue := readFloat("Digite um valor:")
		operation := readOperation()
		otherValue := readFloat("Informe outro valor:")
		result, symbol := calculate(oneValue, otherValue, operation)
		
		fmtResult := fmt.Sprintf("O resultado de %0.2f %s %0.2f é igual á %0.2f", oneValue, symbol, otherValue, result)
		results = append(results, fmtResult)
		fmt.Println(fmtResult)

		fmt.Println("Deseja continuar calculando? (1-Sim, 2-Não)")
		value, _ := reader.ReadString('\n')
		value = regexp.MustCompile(`[+]?\d+`).FindString(value)
		continueCalc , _ := strconv.Atoi(value)
		if continueCalc == 2 {
			break
		}
	}

	fmt.Println("O resultados anteriores são:")
	for _, result := range results {
		fmt.Println(result)
	}


}

func readFloat(label string) float64 {
	for {
		fmt.Println(label)
		value, _ := reader.ReadString('\n')
		value = regexp.MustCompile(`[-+]?\d*\.?\d+`).FindString(value)
		oneValue, err := strconv.ParseFloat(value, 64)
		if (err == nil) {
			return oneValue
		} 
		fmt.Println("O valor informado é inválido", err)
	}
}

func readOperation() int {
	for {
		fmt.Println("Escolha uma operação:")
		fmt.Println("1 - Somar")
		fmt.Println("2 - Subtrair")
		fmt.Println("3 - Multiplicar")
		fmt.Println("4 - Dividir")
		value, _ := reader.ReadString('\n')
		value = regexp.MustCompile(`[+]?\d+`).FindString(value)
		operation, err := strconv.Atoi(value)
		if err == nil {
			return operation
		}
		fmt.Println("A operação informada é inválida", err)
	}
}

func calculate(oneValue, otherValue float64, operation int) (float64, string) {
	var symbol string
	var result float64
	switch operation {
	case 1: //somar
		result = oneValue + otherValue
		symbol = "+"
	case 2: //subtrair
		result = oneValue - otherValue
		symbol = "-"
	case 3: //multiplicar
		result = oneValue * otherValue
		symbol = "*"
	case 4: //dividir
		result = oneValue / otherValue
		symbol = "/"
	}

	return result, symbol
}