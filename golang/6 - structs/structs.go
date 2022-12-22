package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logadouro string
	numero    uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Gil"
	u.idade = 26
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua aqui e agora", 155}

	usuario2 := usuario{"gilvania", 26, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "gilvania"}
	fmt.Println(usuario3)

}
