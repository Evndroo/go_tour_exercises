package main

import (
	"fmt"
	"log"
	"os"

	"github.com/envdroo/learnig-go/geometry"
)

func main() {
	file, err := os.Create("filename.txt")

	if err != nil {
		log.Fatalln("Error on create file:", err)
	}

	file.WriteString("Hello World\n")
	file.WriteString("This is a text that I inserted with Go\n")

	fmt.Println("File created successfully!")

	defer file.Close()

	text, readError := os.ReadFile("filename.txt")

	if readError != nil {
		log.Fatalln("Error on read file:", readError)
	}

	fmt.Println("File content:", string(text))

	fmt.Println("_________________________________________________________")
	fmt.Println("________________ Let's do some Geometry! ________________")
	fmt.Println("_________________________________________________________")
	geometry.Geometry()

	// fmt.Println("_____________________________________________________________________")
	// fmt.Println("___________ And here is my result on the go tour exercise ___________")
	// fmt.Println("__________________________ Uncomment and  ___________________________")
	// fmt.Println("____ Try it on Go Playground or put the image in a base64 viewer ____")
	// fmt.Println("_____________________________________________________________________")
	// geometry.Slices()

	fmt.Println("_____________________________________________________________________")
	fmt.Println("_______ Here is a parsed that I created to exercise maps ____________")
	fmt.Println("_____________________________________________________________________")

	fmt.Println(geometry.ParseStringToWorldMap("I'm Learning Go!"))

	fmt.Println("_____________________________________________________________________")
	fmt.Println("_______ Fibonacci ____________")
	fmt.Println("_____________________________________________________________________")

	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
