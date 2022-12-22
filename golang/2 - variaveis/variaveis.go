package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Outras maneiras de declarar variavel

	var (
		variavel3 string = "tururu"
		variavel4 string = "vai menino"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "Variavel 6"

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	fmt.Println(variavel5, variavel6)

}
