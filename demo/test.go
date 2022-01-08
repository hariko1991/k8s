package main

import "fmt"

func main() {
	ch := make(chan string)

	ch <- "a"
	<-ch

	fmt.Print("ssss")
}
