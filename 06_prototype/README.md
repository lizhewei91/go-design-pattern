# 原型模式
原型模式（Prototype Pattern）是一种创建型设计模式，它允许你通过复制现有对象来创建新对象，而不是通过类构造函数来创建。通过这种方式，你可以避免昂贵的初始化操作，特别是当对象的创建过程非常复杂时。

以下是一个使用 Golang 实现原型模式的示例代码：
```go
package main

import (
	"fmt"
)

// Cloneable 是一个接口，定义了 Clone 方法
type Cloneable interface {
	Clone() Cloneable
}

// 实现了 Cloneable 接口的原型类
type Prototype struct {
	name  string
	value int
}

// Clone 方法创建当前对象的副本
func (p *Prototype) Clone() Cloneable {
	clone := *p
	return &clone
}

func (p *Prototype) GetName() string {
	return p.name
}

func (p *Prototype) GetValue() int {
	return p.value
}

func (p *Prototype) SetName(name string) {
	p.name = name
}

func (p *Prototype) SetValue(value int) {
	p.value = value
}

func main() {
	// 创建一个原型实例
	original := &Prototype{name: "Original", value: 100}

	// 通过克隆原型来创建一个副本
	clone1 := original.Clone().(*Prototype)

	// 修改副本的值
	clone1.SetName("Clone 1")
	clone1.SetValue(200)

	// 打印原型和副本的值
	fmt.Printf("Original: Name = %s, Value = %d\n", original.GetName(), original.GetValue())
	fmt.Printf("Clone 1: Name = %s, Value = %d\n", clone1.GetName(), clone1.GetValue())

	// 通过克隆原型来创建另一个副本
	clone2 := original.Clone().(*Prototype)

	// 打印另一个副本的值
	fmt.Printf("Clone 2: Name = %s, Value = %d\n", clone2.GetName(), clone2.GetValue())
}
```

# 代码说明：

* **Cloneable 接口** ：这个接口定义了一个 Clone 方法，所有需要克隆的对象都应该实现这个接口。

* **Prototype 结构体**：这是一个实现了 Cloneable 接口的结构体，里面包含了两个字段 name 和 age，以及一些方法用来获取和设置这些字段。

* **Clone 方法**：这个方法通过值拷贝的方式创建了一个当前对象的副本，并返回一个 Cloneable 接口类型的对象。

* **main 函数**：在 main 函数中，首先创建了一个 Prototype 的实例 original，然后通过调用 Clone 方法创建了它的两个副本，并对第一个副本进行了修改，最后打印出原型和副本的值。

# 输出结果：
```shell
Original: Name = Original, Value = 100
Clone 1: Name = Clone 1, Value = 200
Clone 2: Name = Original, Value = 100
```

在这个例子中，克隆操作使得 Clone 1 与 Original 相互独立地存在和操作，改变 Clone 1 的值不会影响到 Original。