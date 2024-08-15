package builder

import "fmt"

// PersonInterface 接口表示一个人的信息和行为
type PersonInterface interface {
	Speak(message string)
	Sleep()
	GetName() string
}

// Person 结构体实现了 PersonInterface 接口
type Person struct {
	name    string
	age     string
	gender  string
	address string
}

// Speak 方法实现了 PersonInterface 接口中的说话方法
func (p *Person) Speak(message string) {
	fmt.Printf("%s says: %s\n", p.name, message)
}

// Sleep 方法实现了 PersonInterface 接口中的睡觉方法
func (p *Person) Sleep() {
	fmt.Printf("%s is sleeling now\n", p.name)
}

// GetName 方法实现了 PersonInterface 接口中的获取名称方法
func (p *Person) GetName() string {
	return p.name
}

// PersonBuilderInterface 接口表示一个人的信息构造器
type PersonBuilderInterface interface {
	SetName(name string) PersonBuilderInterface
	SetAge(age string) PersonBuilderInterface
	SetGender(gender string) PersonBuilderInterface
	SetAddress(address string) PersonBuilderInterface
	Build() PersonInterface
}

// PersonBuilder 结构体表示一个人的建造器
type PersonBuilder struct {
	name    string
	age     string
	gender  string
	address string
}

// SetName 方法设置人的姓名
func (b *PersonBuilder) SetName(name string) PersonBuilderInterface {
	b.name = name
	return b
}

// SetAge 方法设置人的年龄
func (b *PersonBuilder) SetAge(age string) PersonBuilderInterface {
	b.age = age
	return b
}

// SetGender 方法设置人的性别
func (b *PersonBuilder) SetGender(gender string) PersonBuilderInterface {
	b.gender = gender
	return b
}

// SetAddress 方法设置人的地址
func (b *PersonBuilder) SetAddress(address string) PersonBuilderInterface {
	b.address = address
	return b
}

// Build 方法创建一个新的 Person 实例
func (b *PersonBuilder) Build() PersonInterface {
	return &Person{
		name:    b.name,
		age:     b.age,
		gender:  b.gender,
		address: b.address,
	}
}
