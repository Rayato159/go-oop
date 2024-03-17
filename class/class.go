package main

import "fmt"

type Airplane struct {
	Name string
}

func (a Airplane) Fly() {
	fmt.Printf("%s: Fly\n", a.Name)
}

func main() {
	spitfire := Airplane{Name: "Spitfire"}
	boeing := Airplane{Name: "Boeing"}

	spitfire.Fly()
	boeing.Fly()
}
