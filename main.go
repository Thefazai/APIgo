package main

import (
	"fmt"
	"os"

	"apigo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:, err")
		os.Exit(1)
	}
}
