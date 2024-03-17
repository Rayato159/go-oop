package main

import "fmt"

type (
	Keyboard interface {
		Input()
	}

	MechanicalKeyboard struct {
		SwitchType string
		Size       string
		OS         string
	}

	NormalKeyboard struct {
		IsCanPress bool
	}
)

func (m MechanicalKeyboard) Input() {
	fmt.Println("Pressing the key on the mechanical keyboard: ", m.SwitchType, m.Size, m.OS)
}

func (m NormalKeyboard) Input() {
	fmt.Println("Pressing the key on the keyboard")
}

func main() {
	mechanicalKeyboard := MechanicalKeyboard{
		SwitchType: "Akko Cream Yellow Magnetic Switch",
		Size:       "75%",
		OS:         "Windows,MacOS",
	}

	normalKeyboard := NormalKeyboard{
		IsCanPress: true,
	}

	mechanicalKeyboard.Input()
	normalKeyboard.Input()
}
