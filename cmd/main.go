package main

import (
	"fmt"
	"log"
	"os"
	"stocks_tracker/service"
)

func main(){
	log.Printf("** booting stocks-tracker service **")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	return service.New()
}
