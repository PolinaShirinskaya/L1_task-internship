// Реализовать паттерн «адаптер» на любом примере

package main

import "fmt"

// Transport - интерфейс транспорта
type Transport interface {
	Drive()
}

// Auto - класс машины
type Auto struct{}

func (a Auto) Drive() {
	fmt.Println("Машина едет по дороге")
}

// Driver - класс путешественника
type Driver struct{}

func (d Driver) Travel(transport Transport) {
	transport.Drive()
}

// Animal - интерфейс животного
type Animal interface {
	Move()
}

// Camel - класс верблюда
type Camel struct{}

func (c Camel) Move() {
	fmt.Println("Верблюд идет по пескам пустыни")
}

// CamelToTransportAdapter - адаптер от Camel к Transport
type CamelToTransportAdapter struct {
	camel Camel
}

func (a CamelToTransportAdapter) Drive() {
	a.camel.Move()
}

func main() {
	// путешественник
	driver := Driver{}
	// машина
	auto := Auto{}
	// отправляемся в путешествие
	driver.Travel(auto)

	// встретились пески, надо использовать верблюда
	camel := Camel{}
	// используем адаптер
	camelTransport := CamelToTransportAdapter{camel}
	// продолжаем путь по пескам пустыни
	driver.Travel(camelTransport)
}
