package main

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

type Prototype struct {
	name string
	age  int
}

func (p *Prototype) Clone() Cloneable {
	clone := *p
	return &clone
}

func (p *Prototype) SetName(name string) {
	p.name = name
}

func (p *Prototype) SetAge(age int) {
	p.age = age
}

func (p *Prototype) GetName() string {
	return p.name
}

func (p *Prototype) GetAge() int {
	return p.age
}

func main() {
	original := &Prototype{name: "lzw", age: 18}

	clone1 := original.Clone().(*Prototype)

	clone1.SetName("lx")
	clone1.SetAge(19)
	fmt.Printf("Original name:%v age:%v\n", original.GetName(), original.GetAge())
	fmt.Printf("clone1 name:%v age:%v\n", clone1.GetName(), clone1.GetAge())

	clone2 := original.Clone().(*Prototype)
	fmt.Printf("clone2 name:%v age:%v\n", clone2.GetName(), clone2.GetAge())
}
