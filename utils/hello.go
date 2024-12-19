package utils

import "fmt"

func DisplayIntroduction() {
	asciiArt := `
   ______         ______          __  _             __
  / ____/___     /_  __/__  _____/ /_(_)___  ____ _/ /
 / / __/ __ \     / / / _ \/ ___/ __/ / __ \/ __ '/ / 
/ /_/ / /_/ /    / / /  __(__  ) /_/ / / / / /_/ /_/  
\____/\____/    /_/  \___/____/\__/_/_/ /_/\__, (_)   
                                          /____/      `

	fmt.Println(asciiArt)
	fmt.Printf("Bem vindo ao compilador de testes! Aqui há vários scripts feitos por JVVN para testar a linguagem Go!\nEscolha o que inicializar:\n")
	fmt.Print("#> ")

}

func HelloWorld() {
	fmt.Println("Hello, World!")
}
