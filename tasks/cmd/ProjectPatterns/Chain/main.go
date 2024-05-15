// Паттерн предлагает связать объекты обработчиков в одну цепь. Каждый из них будет иметь ссылку на следующий обработчик в цепи. Таким образом, при получении запроса обработчик сможет не только сам что-то с ним сделать, но и передать обработку следующему объекту в цепочке.
package main

import "fmt"

// Request представляет запрос пользователя.
type Request struct {
	amount int
}

// Handler определяет интерфейс для обработчика запроса.
type Handler interface {
	SetNext(handler Handler)
	Handle(request *Request)
}

// BaseHandler предоставляет базовую реализацию обработчика.
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// ConcreteHandler представляет конкретного обработчика.
type ConcreteHandler struct {
	BaseHandler
}

func (h *ConcreteHandler) Handle(request *Request) {
	if request.amount < 1000 {
		fmt.Println("Запрос обработан обработчиком 1")
	} else if h.next != nil {
		h.next.Handle(request)
	} else {
		fmt.Println("Запрос не может быть обработан")
	}
}

func main() {
	// Создание обработчиков
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}
	handler3 := &ConcreteHandler{}

	// Установка цепочки обработчиков
	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	// Обработка запросов разного размера
	request1 := &Request{amount: 500}
	handler1.Handle(request1)

	request2 := &Request{amount: 1500}
	handler1.Handle(request2)
}

// Request представляет запрос пользователя.
// Handler определяет интерфейс для обработчика запроса, а BaseHandler предоставляет базовую реализацию.
// ConcreteHandler представляет конкретного обработчика, который может либо обработать запрос, либо передать его следующему обработчику в цепочке.
// В функции main создаются обработчики и устанавливается цепочка обработчиков. Затем обработчики вызываются для обработки запросов разного размера.
// Паттерн "цепочка обязанностей" полезен, когда у вас есть несколько объектов, которые могут обработать запрос, и вы хотите, чтобы запрос автоматически передавался от одного объекта к другому в цепочке, пока не будет найден объект, способный обработать запрос. В нашем примере, если запрос не может быть обработан первым обработчиком, он передается следующему, и так далее, пока не будет найден подходящий обработчик.
