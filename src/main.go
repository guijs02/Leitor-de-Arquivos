package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Hello World!")
	leHeroisDoArquivo()
	
}

func leHeroisDoArquivo (){
	const arquivoHerois string = "herois.txt"
	arquivo, err := os.ReadFile(arquivoHerois)
	
	if err != nil{
		fmt.Println("Erro ao abrir o arquivo", err)
		return
	}

	fmt.Println(string(arquivo))
}
