package pattern

import (
	"fmt"
)

/* Паттерн фабрика позволяет ввести новый уровень абстракции, сокрыть реализацию объектов и ввести некоторую инкапсуляцию в приложение.
С помощью него вы данном примере мы можем создать логгер одного из двух тиипов, используя интерфейсы Logger и LoggerFactory.

В реальных примерах такой паттерн пригодится для сложных объектов, дабы упростить их создание для читабельности кода.

К плюсам можно отнести хорошую читаемость и удобство использования в дальнейшем коде.

К минусам иногда можно отнести излишнюю инкапсуляцию.
*/

type Logger interface {
	Log(message string)
}

type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	fmt.Println("Logging to file:", message)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Logging to console:", message)
}

type LoggerFactory interface {
	CreateLogger() Logger
}

type FileLoggerFactory struct{}

func (f *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

type ConsoleLoggerFactory struct{}

func (c *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}

func FactoryMain() {
	fileLoggerFactory := FileLoggerFactory{}
	fileLogger := fileLoggerFactory.CreateLogger()
	fileLogger.Log("Hello from File logger")

	consoleLoggerFactory := ConsoleLoggerFactory{}
	consoleLogger := consoleLoggerFactory.CreateLogger()
	consoleLogger.Log("Hello from Console logger")
}
