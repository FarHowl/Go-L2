package pattern

import (
	"fmt"
	"math"
)

/*
Паттерн Посетитель позволяет добавлять дополнительную функциональность для уже существующих объектов.
В рамках его применения необходимо добавить лишь один метод для оригинального объекта.
Далее с помощью интерфейса Visitor можно определять конкретных посетителей, который будут реализовывать методы посещения всех наших объектов.
Все взаимодействие происходит через метод accept(), который принимает конкретного посетителя.
В данном методе происходит вызов метода visit для конкретного объекта.

В качестве примера реализации такого паттерна на практике можно привести внедрение новой функциональности в библиотеки.
Внедение одного нового метода вряд ли сможет затронуть уже существующую функциональность. А затем можно добавлять любое количество
посетителей.

К плюсам можно отнести относительную полезность при поддержке уже существующих библиотек.

К минусам можно отнести то, что в некоторых случаях придется добавлять действительно много кода для внесения новой функциональности.
*/
type Shape interface {
	accept(Visitor)
}

type Visitor interface {
	visitForRectangle(*Rectangle) int
	visitForSquare(*Square) int
	visitForCircle(*Circle) float64
}

type Rectangle struct {
	a, b int
}

func (r *Rectangle) accept(v Visitor) int {
	return v.visitForRectangle(r)
}

type Circle struct {
	r int
}

func (c *Circle) accept(v Visitor) float64 {
	return v.visitForCircle(c)
}

type Square struct {
	a int
}

func (s *Square) accept(v Visitor) int {
	return v.visitForSquare(s)
}

type AreaCalculator struct {
}

func (a *AreaCalculator) visitForCircle(c *Circle) float64 {
	return math.Pi * float64(c.r) * float64(c.r)
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) int {
	return r.a * r.b
}

func (a *AreaCalculator) visitForSquare(s *Square) int {
	return s.a * s.a
}

func VisitorMain() {
	rect := Rectangle{2, 5}
	circ := Circle{4}
	sqr := Square{3}

	areaCalculator := AreaCalculator{}

	circleArea := circ.accept(&areaCalculator)
	squareArea := sqr.accept(&areaCalculator)
	rectangleArea := rect.accept(&areaCalculator)

	fmt.Printf("Circle area : %f", circleArea)
	fmt.Printf("\nSquare area : %d", squareArea)
	fmt.Printf("\nRectangle area : %d", rectangleArea)
}
