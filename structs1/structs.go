package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a Animal) Eat() {
	fmt.Println(a.food)
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{food: "grass", locomotion: "walk", sound: "moo"},
		"snake": Animal{food: "mice", locomotion: "slither", sound: "hssss"},
		"bird":  Animal{food: "worms", locomotion: "fly", sound: "peep"},
	}

	for {
		var animalName, command string
		fmt.Print("> ")
		fmt.Scanf("%s %s", &animalName, &command)

		animal, ok := animals[animalName]
		if !ok {
			panic("this animal doesn't exist")
		}

		if command == "eat" {
			animal.Eat()
		} else if command == "speak" {
			animal.Speak()
		} else if command == "move" {
			animal.Move()
		} else {
			panic("Unknown command")
		}
	}
}
