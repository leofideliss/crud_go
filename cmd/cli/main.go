package main

import (
	"fmt"
	"os"    
	"crud_go/cmd/cli/commands"
)

func main() {    
	if err := commands.Execute(); err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}
