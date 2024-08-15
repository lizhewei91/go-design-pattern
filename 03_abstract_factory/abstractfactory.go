package abstractfactory

import "fmt"

// 抽象工厂
type ISportFactory interface {
	makeShoe() IShoe
}

func NewSportFactory(brand string) (ISportFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type padded")
}

// 具体工厂
type Adidas struct{}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

type Nike struct{}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 16,
		},
	}
}

// 抽象产品
type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

// 具体产品
type AdidasShoe struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}
