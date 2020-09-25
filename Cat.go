package model

import "fmt"

type Cat struct {
	name string
}

func NewCat(name string) *Cat {
	var c Cat
	c.name = name
	return &c
}

func (c *Cat) Sounds() {
	fmt.Printf("%s: meow.\n", c.name)
} //클래스
