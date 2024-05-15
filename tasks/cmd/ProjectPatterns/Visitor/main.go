package main

import "fmt"

// Visitor определяет интерфейс для посетителя.
type Visitor interface {
	VisitHome(*Home)
	VisitBank(*Bank)
	VisitFactory(*Factory)
}

// Element определяет интерфейс для элемента, который принимает посетителя.
type Element interface {
	Accept(Visitor)
}

// Home представляет дом.
type Home struct{}

func (h *Home) Accept(visitor Visitor) {
	visitor.VisitHome(h)
}

// Bank представляет банк.
type Bank struct{}

func (b *Bank) Accept(visitor Visitor) {
	visitor.VisitBank(b)
}

// Factory представляет фабрику.
type Factory struct{}

func (f *Factory) Accept(visitor Visitor) {
	visitor.VisitFactory(f)
}

// InsuranceAgent представляет страхового агента.
type InsuranceAgent struct{}

func (ia *InsuranceAgent) VisitHome(home *Home) {
	fmt.Println("Страховой агент предлагает медицинскую страховку для дома")
}

func (ia *InsuranceAgent) VisitBank(bank *Bank) {
	fmt.Println("Страховой агент предлагает страховку от грабежа для банка")
}

func (ia *InsuranceAgent) VisitFactory(factory *Factory) {
	fmt.Println("Страховой агент предлагает страховку от пожара и наводнения для фабрики")
}

func main() {
	agent := &InsuranceAgent{}

	// Посещение различных мест:
	home := &Home{}
	bank := &Bank{}
	factory := &Factory{}

	// Принятие страхового агента на различных местах:
	home.Accept(agent)
	bank.Accept(agent)
	factory.Accept(agent)
}

// Visitor определяет интерфейс для посетителя, который может посещать различные типы объектов.
// Element определяет интерфейс для элемента, который может принимать посетителя.
// Home, Bank и Factory представляют различные типы мест.
// InsuranceAgent представляет страхового агента и реализует методы посещения для каждого типа места.
// Паттерн "посетитель" полезен, когда у вас есть структура объектов с разными типами, и вы хотите добавить новые операции без изменения этих объектов. В нашем примере, страховой агент (посетитель) может посещать различные места и предлагать соответствующие страховые продукты для каждого места.
