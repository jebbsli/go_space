package main

import (
	"flag"
	"log"
)

func main() {
	var name string

	flag.StringVar(&name, "name", "Golang", "help info")
	flag.StringVar(&name, "n", "Golang", "help info")
	flag.Parse()

	log.Printf("Name: %s", name)
}
