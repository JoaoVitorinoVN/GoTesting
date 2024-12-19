package main

import (
	"fmt"
	"go-testing/utils"
	"time"
)

func main() {
	var userChoice int8
	utils.DisplayIntroduction()

	fmt.Scan(&userChoice)
	utils.Execute(userChoice)

	fmt.Println("\n\nObrigado por usar o Go Testing!")
	time.Sleep(3 * time.Second)
}
