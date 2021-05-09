package main

import (
	"fmt"
	"os"
	"stocks_tracker/service"
)

func main(){
	fmt.Printf("** booting stocks-tracker service **\n")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	return service.New()
}
