package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("caminho/do/arquivo.txt")
	if err != nil {
		// Lida com o erro, se ocorrer
	}
	defer file.Close() // Certifique-se de fechar o arquivo quando terminar

	// Faça ações adicionais no arquivo, como leitura de conteúdo
	scanner := bufio.NewScanner(file)
	var content string

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		// Lida com o erro, se ocorrer
	}
}
