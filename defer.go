package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runApplication()
}
