package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "Your Filename.txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Um erro ocorreu: ", err)
	} else if err != nil {
		fmt.Println("Nenhum erro, operação com sucesso")
	}
	content := "Your Content Here"
	_, err = file.WriteString(content)
	if err != nil {

	}
	defer file.Close()
}
