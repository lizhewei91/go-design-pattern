package factorymethod

// 产品方法
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// 具体产品
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

type AK47 struct {
	Gun
}

func NewAk47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "ak47 gun",
			power: 4,
		},
	}
}

type musket struct {
	Gun
}

func NewMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "musket gun",
			power: 1,
		},
	}
}

// 工厂方法
func NewGunType(gunType string) IGun {
	switch gunType {
	case "ak47":
		return NewAk47()
	case "musket":
		return NewMusket()
	}
	return nil
}
