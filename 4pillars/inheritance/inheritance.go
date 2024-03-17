package main

import "fmt"

type (
	Human struct{}

	Devil struct {
		Human
	}

	Elf struct {
		Human
	}
)

func (h Human) Walk() {
	fmt.Println("Walking")
}

func (h Human) Eat() {
	fmt.Println("Eating")
}

func (h Human) Breath() {
	fmt.Println("Breathing")
}

func (d Devil) Mutate() {
	fmt.Println("Mutating")
}

func (d Devil) Fly() {
	fmt.Println("Flying")
}

func (e Elf) MagicSpell() {
	fmt.Println("Casting a magic spell")
}

func main() {
	human := Human{}
	devil := Devil{}
	elf := Elf{}

	fmt.Println("Human: ")
	human.Walk()
	human.Eat()
	human.Breath()
	fmt.Println()

	fmt.Println("Devil: ")
	devil.Walk()
	devil.Eat()
	devil.Breath()
	devil.Mutate()
	devil.Fly()
	fmt.Println()

	fmt.Println("Elf: ")
	elf.Walk()
	elf.Eat()
	elf.Breath()
	elf.MagicSpell()
}
