package main

import (
	"fmt"
	"log"
	"manish1457/test/project-test/pkg"
)

func main() {
	log.Println("project-test")
	log.SetPrefix("pkg: ")
	log.SetFlags(0)
	// request a greeting messege
	message, err := pkg.Name("Manish")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	//a slice of names
	names := []string{"Manish", "Rahul", "Jack", "Umesh"}

	messages, err := pkg.Names(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
