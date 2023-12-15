package pattern

import "fmt"

/* Паттерн Команда позволяет инкапсулировать запросы от инициатора к получателю, добавлять их в очередь и так далее.
В данном примере был реализован получатель в виде светильника и инициатор запроса на выполнение команд светильника в виде
пульта ДУ.

В реальных примерах данный паттерн может быть использован для создания механизма транзакций, очереди запросов.

К плюсам можно отнести создаваемую инкапсуляцию вызова определенных методов получателя.const

К минусам можно отнести избыточность кода в некоторых случаях.
*/

type Command interface {
	Execute()
}

type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
}

func (l *Light) TurnOff() {
	l.isOn = false
}

type TurnOnCommand struct {
	Light *Light
}

func (t *TurnOnCommand) Execute() {
	t.Light.TurnOn()
}

type TurnOffCommand struct {
	Light *Light
}

func (t *TurnOffCommand) Execute() {
	t.Light.TurnOff()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) setCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) pressButton() {
	r.command.Execute()
}

func CommandMain() {
	light := Light{}
	turnOnCommand := TurnOnCommand{Light: &light}
	turnOffCommand := TurnOffCommand{Light: &light}

	remoteControl := RemoteControl{}
	remoteControl.setCommand(&turnOffCommand)
	remoteControl.pressButton()
	fmt.Println(light.isOn)

	remoteControl.setCommand(&turnOnCommand)
	remoteControl.pressButton()
	fmt.Println(light.isOn)
}
