package main

import (
	"comand-project/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting project")
	aplication := app.Gerar()

	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
