package main

import (
	"fmt"
	"gstore/service"
)

func main() {
	fmt.Println("train: ")
	service.Train()
	fmt.Println("\ntest: ")
	service.Test()
}
