package main

import "fmt"

type Animal struct {
	Limbs int
	Sound string
}

func (a Animal) MakeNoise() string {
	return a.Sound
}

type Dog struct {
	Animal
}

func main() {
	d1 := Dog{
		Animal{
			Limbs: 4,
			Sound: "bow",
		},
	}

	fmt.Println("Dog: ",d1.MakeNoise())

	fmt.Println(d1.Limbs, d1.Sound)
}
