/*
Copyright © 2024 HSINE KUEI CHIU <theone1632@gmail.com>
*/
package main

import (
	"log"
	"os"

	"github.com/brionac626/api-demo/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println("Execute command failed", err)
		os.Exit(1)
	}
}
