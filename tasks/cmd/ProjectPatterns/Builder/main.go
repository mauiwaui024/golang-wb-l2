package main

import "fmt"

// Product представляет конечный объект, который мы строим.
type Product struct {
	part1 string
	part2 string
	part3 string
}

// Builder определяет интерфейс для создания различных частей продукта.
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetProduct() *Product
}

// ConcreteBuilder представляет конкретную реализацию Builder.
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{product: &Product{}}
}

func (cb *ConcreteBuilder) BuildPart1() {
	cb.product.part1 = "Part1 built"
}

func (cb *ConcreteBuilder) BuildPart2() {
	cb.product.part2 = "Part2 built"
}

func (cb *ConcreteBuilder) BuildPart3() {
	cb.product.part3 = "Part3 built"
}

func (cb *ConcreteBuilder) GetProduct() *Product {
	return cb.product
}

// Director определяет процесс сборки продукта с использованием Builder.
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() *Product {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
	return d.builder.GetProduct()
}

func main() {
	builder := NewConcreteBuilder()
	director := NewDirector(builder)

	product := director.Construct()

	fmt.Println("Product built:")
	fmt.Println("Part 1:", product.part1)
	fmt.Println("Part 2:", product.part2)
	fmt.Println("Part 3:", product.part3)
}

// Product представляет конечный объект, который мы строим.
// Builder определяет интерфейс для создания различных частей продукта.
// ConcreteBuilder представляет конкретную реализацию Builder. Он заботится о создании конкретных частей продукта.
// Director определяет процесс сборки продукта с использованием Builder. Он скрывает сложность сборки продукта от клиента.
