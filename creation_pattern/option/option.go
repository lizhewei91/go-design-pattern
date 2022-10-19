package main

type Person struct {
	Name string  // 固定参数name
	Age  int     // 固定参数age
	Opt  options // 可变参数
}

type options struct {
	location string
	level    string
}

type OptionFunc func(*options)

func NewPerson(name string, age int, opts ...OptionFunc) *Person {
	p := &Person{
		Name: name,
		Age:  age,
		Opt:  defaultOptions,
	}

	for _, opt := range opts {
		opt(&p.Opt)
	}
	return p
}

// 可变参数，设置默认参数
var defaultOptions = options{
	location: "beijing",
	level:    "collage",
}

func WithLocation(location string) OptionFunc {
	return func(opt *options) {
		opt.location = location
	}
}

func WithLevel(level string) OptionFunc {
	return func(opt *options) {
		opt.level = level
	}
}

func Run(name, location, level string, age int) *Person {
	p := NewPerson(
		name,
		age,
		WithLocation(location),
		WithLevel(level),
	)
	return p
}
