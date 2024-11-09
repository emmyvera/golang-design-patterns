package main

import "fmt"

// Animal is the type for our Abstact factory
type Animal interface {
	Says()
	LikeWater() bool
}

// Dog the concrete factory for dogs
type Dog struct{}

func (d *Dog) Says() {
	fmt.Println("Woof")
}

func (d *Dog) LikeWater() bool {
	return true
}

// Cat The concrete factory for cats
type Cat struct{}

func (c *Cat) Says() {
	fmt.Println("Meow")
}

func (c *Cat) LikeWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct{}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	catFactory := CatFactory{}
	dogFactory := DogFactory{}

	cat := catFactory.New()
	dog := dogFactory.New()

	cat.Says()
	dog.Says()

	fmt.Println("A dog likes water: ", dog.LikeWater())
	fmt.Println("A cat likes water: ", cat.LikeWater())

}
