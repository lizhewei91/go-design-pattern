package singleton

import "sync"

// Singleton 是单例模式接口，导出的
// 通过该接口可以避免 GetInstance 返回一个包私有类型的指针
type Singleton interface {
	GetName() string
	GetAge() int
}

// singleton 结构体定义了一个单例对象
type singleton struct {
	name string
	age  int
}

func (s *singleton) GetName() string {
	return s.name
}

func (s *singleton) GetAge() int {
	return s.age
}

// instance 是一个包级别的变量，用于保存唯一的单例对象
var instance *singleton

// once 是一个用于确保初始化操作只执行一次的 sync.Once 对象
var once sync.Once

// GetInstance 用于获取单例模式对象
// 第一次调用时初始化单例，并允许传入参数 name 和 age
func GetInstance(name string, age int) Singleton {
	// 使用 sync.Once 确保初始化操作只执行一次
	once.Do(func() {
		instance = &singleton{
			name: name,
			age:  age,
		}
	})
	return instance
}
