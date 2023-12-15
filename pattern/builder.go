package pattern

import "fmt"

/* Паттерн строитель повзоляет создавать сложные объекты разных спецификаций. Поля объектов одинаковые, но значения - разные.
Он позволяет избежать передачи множества аргументов, используя ряд методов, каждый из которых задает определенные поля объекта.
С помощью него мы можем создаввать разные объекты, используя один и тот же процесс.

В реальных примерах его можно использовать там, где для создания объектов требуется вызывать дополнительные методы, логгеры
и другие для полей объекта

К плюсам можно отнести сокрытие реализации больших конструкторов

К минусам хочется отнести то, что его можно использовать далеко не везде. Иногда он будет избыточным, занимать много места
*/

type Car struct {
	Model    string
	Engine   string
	Interior string
	Wheels   string
}

type CarBuilder interface {
	BuildModel()
	BuildEngine()
	BuildInterior()
	BuildWheels()
	getCar() Car
}

type SportCarBuilder struct {
	Car
}

func (s *SportCarBuilder) BuildModel() {
	s.Car.Model = "SportCar model"
}

func (s *SportCarBuilder) BuildEngine() {
	s.Car.Engine = "SportCar engine"
}

func (s *SportCarBuilder) BuildInterior() {
	s.Car.Interior = "SportCar interior"
}

func (s *SportCarBuilder) BuildWheels() {
	s.Car.Wheels = "SportCar wheels"
}

func (s *SportCarBuilder) getCar() Car {
	return s.Car
}

type SedanCarBuilder struct {
	Car
}

func (s *SedanCarBuilder) BuildModel() {
	s.Car.Model = "SedanCar model"
}

func (s *SedanCarBuilder) BuildEngine() {
	s.Car.Engine = "SedanCar engine"
}

func (s *SedanCarBuilder) BuildInterior() {
	s.Car.Interior = "SedanCar interior"
}

func (s *SedanCarBuilder) BuildWheels() {
	s.Car.Wheels = "SedanCar wheels"
}
func (s *SedanCarBuilder) getCar() Car {
	return s.Car
}

type Director struct {
	builder CarBuilder
}

func (d *Director) BuildCar(builder CarBuilder) {
	d.builder = builder
	builder.BuildEngine()
	builder.BuildInterior()
	builder.BuildModel()
	builder.BuildWheels()
}

func BuilderMain() {
	director := Director{}
	sportCarBuilder := SportCarBuilder{}
	sedanCarBuilder := SedanCarBuilder{}

	director.BuildCar(&sportCarBuilder)
	director.BuildCar(&sedanCarBuilder)

	fmt.Println(sportCarBuilder.getCar())
	fmt.Println(sedanCarBuilder.getCar())
}
