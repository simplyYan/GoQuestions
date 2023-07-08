package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 3 {
			fmt.Println("Está na metade")
		}
	}

}
