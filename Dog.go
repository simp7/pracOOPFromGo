package main

import "fmt"

type Dog struct {
	name string
}

func NewDog(name string) Dog {
	var d Dog
	d.name = name
	return d
} // 구조체

func (d Dog) Sounds() {
	fmt.Printf("%s: woof woof!\n", d.name)

}