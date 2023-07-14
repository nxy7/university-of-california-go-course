package main

import (
	"fmt"
)

type Animal interface {
	Speak()
	Move()
	Eat()
}

type Snake struct {
	name string
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}
func (s Snake) Move() {
	fmt.Println("sligther")
}
func (s Snake) Eat() {
	fmt.Println("mice")
}

type Cow struct {
	name string
}

func (c Cow) Speak() {
	fmt.Println("mooo")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Eat() {
	fmt.Println("grass")
}

type Bird struct {
	name string
}

func (b Bird) Speak() {
	fmt.Println("peep")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Eat() {
	fmt.Println("worms")
}

func main() {
	animals := map[string]Animal{}

	for {
		var mainCommand, arg1, arg2 string
		fmt.Print("> ")
		fmt.Scanf("%s %s %s", &mainCommand, &arg1, &arg2)

		if mainCommand == "query" {
			QueryAnimal(animals, arg1, arg2)
		} else if mainCommand == "newanimal" {
			AddAnimal(animals, arg1, arg2)
		} else {
			panic("Unknown command")
		}

	}
}

func QueryAnimal(animals map[string]Animal, animalName, command string) {
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

func AddAnimal(animals map[string]Animal, animalName, kind string) {
	_, ok := animals[animalName]
	if ok {
		panic("this animal already exists")
	}

	if kind == "cow" {
		animals[animalName] = Cow{name: animalName}
	} else if kind == "bird" {
		animals[animalName] = Bird{name: animalName}

	} else if kind == "snake" {
		animals[animalName] = Snake{name: animalName}
	} else {
		panic("Unknown animal kind")
	}
}
