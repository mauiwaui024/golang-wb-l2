package main

// Subsystem1 и Subsystem2 представляют сложные подсистемы со своим набором операций.
// Facade предоставляет упрощенный интерфейс к этим подсистемам. Он скрывает сложность взаимодействия с подсистемами от клиента.
// Функция main демонстрирует, как использовать фасад для выполнения операций с подсистемами без необходимости знать их внутренние детали.
// Этот шаблон полезен, когда вы хотите предоставить простой интерфейс к сложной системе или когда вам нужно отделить код клиента от подсистем.

import "fmt"

// Subsystem1 представляет собой сложную подсистему.
type Subsystem1 struct{}

func (s *Subsystem1) Operation1() {
	fmt.Println("Подсистема 1: Какая-то сложная операция 1")
}

func (s *Subsystem1) Operation2() {
	fmt.Println("Подсистема 1: Какая-то сложная операция 2")
}

// Subsystem2 представляет другую сложную подсистему.
type Subsystem2 struct{}

func (s *Subsystem2) Operation3() {
	fmt.Println("Подсистема 2: Какая-то сложная операция 3")
}

func (s *Subsystem2) Operation4() {
	fmt.Println("Подсистема 2: Какая-то сложная операция 4")
}

// Facade предоставляет простой интерфейс для взаимодействия с комплексной подсистемой.
type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
}

func NewFacade() *Facade {
	return &Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
	}
}

func (f *Facade) Operation() {
	fmt.Println("Фасад: операция")
	f.subsystem1.Operation1()
	f.subsystem1.Operation2()
	f.subsystem2.Operation3()
	f.subsystem2.Operation4()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
