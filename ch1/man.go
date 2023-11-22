package main

import "fmt"

type Man struct {
	Name string
}

func NewMan(name string) Man {
	man := Man{name}
	fmt.Println("Initialized!")
	return man
}

func (m Man) Hello() {
	fmt.Printf("Hell %s!\n", m.Name)
}

func (m Man) Goodbye() {
	fmt.Printf("Good-bye %s!\n", m.Name)
}
