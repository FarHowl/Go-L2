package pattern

import "fmt"

/* Паттерн Состояние позволяет хранить в объектах состояние и проверять его, не используя дополнительные условные конструкции.
Для этого обычно используется интерфейс состояния, который могут реализовать несколько конкретных состояний.

В реальных примерах такое может использоваться в сложных объектах, дабы избежать многократной проверки состояния в выполняющейся программе.

К плюсам можно отнести действитетельно удобную конструкцию и простую реализацию.

К минусам можно отнести небольшие, по сравнению с предыдущими паттернами, расходы.
*/

type Switch struct {
	state State
}

func (sw *Switch) SetState(state State) {
	sw.state = state
}

func (sw *Switch) On() {
	sw.state.On(sw)
}

func (sw *Switch) Off() {
	sw.state.Off(sw)
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type OnState struct{}

func (o *OnState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning light off")
	sw.SetState(&OffState{})
}

type OffState struct{}

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on")
	sw.SetState(&OnState{})
}

func (o *OffState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

func StateMain() {
	sw := Switch{&OnState{}}
	sw.On()
	sw.Off()
	sw.Off()
}
