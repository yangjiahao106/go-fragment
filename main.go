package main

import "fmt"

func main() {
	fmt.Println("hello world")
	{
		foo := NewFoo(WithName("name"), WithAge(1), WithDB("db"))
		fmt.Println(foo)

	}
	foo := "xx"
	fmt.Println(foo)
}

type Foo struct {
	name string
	age  int
	db   interface{}
}

type Option func(foo *Foo)

func WithName(name string) Option {
	return func(foo *Foo) {
		foo.name = name
	}
}

func WithAge(age int) Option {
	return func(foo *Foo) {
		foo.age = age
	}
}

func WithDB(db interface{}) Option {
	return func(foo *Foo) {
		foo.db = db
	}
}

func NewFoo(options ...Option) *Foo {
	foo := &Foo{}
	for _, op := range options {
		op(foo)
	}
	return foo
}
