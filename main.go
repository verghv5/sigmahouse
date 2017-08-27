package main

//import "os"

/*
 main file used to initialize server
 */

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"root",
		"postgres")

	a.Run(":8000")
}