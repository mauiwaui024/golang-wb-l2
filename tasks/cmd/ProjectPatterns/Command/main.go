package main

import "fmt"

// Command определяет интерфейс для выполнения команды.
type Command interface {
	Execute()
}

// Receiver представляет получателя команды.
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Повар готовит блюда по заказу...")
}

// ConcreteCommand представляет конкретную команду.
type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

// Invoker представляет вызывающего.
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	// Создание получателя
	receiver := &Receiver{}

	// Создание команды с указанием получателя
	command := NewConcreteCommand(receiver)

	// Создание вызывающего и назначение ему команды
	invoker := &Invoker{}
	invoker.SetCommand(command)

	// Вызов команды
	invoker.ExecuteCommand()
}

// Command определяет интерфейс для выполнения команды.
// Receiver представляет получателя команды, который знает, как выполнять действие.
// ConcreteCommand представляет конкретную команду, которая связывает получателя и действие.
// Invoker представляет вызывающего, который запускает команду.
// В нашем сценарии, официант (Invoker) принимает заказ и передает команду (заказ) повару (Receiver), который затем выполняет необходимое действие (готовит блюда). При этом официант не знает, как именно повар выполняет заказ, а повар не знает, кто отправил заказ. Таким образом, обе стороны остаются изолированными друг от друга.
