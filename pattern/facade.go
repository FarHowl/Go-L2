package pattern

import "fmt"

/*В данном примере мы используем паттерн фасад для единой точки взаимодействия с несколькими системами.
Внутри каждой системы могут быть долгие и сложные вызовы различных функций, однако наша структура FacadeAPI
берет на себя вызов всех внутренних API для получения определенных данных. На выходе мы получаем лишь один метод.

В реальных примерах фасад так же используется для реализации API Gateway, который предоставляет единый доступ ко всем внутренним API,
скрывая их реализацию.

К плюсам можно отнести простоту использования готового интерфейса, сокрытие сложности реализации подсистем.
К минусам можно отнести новый слой абстракции, который создает фасад, а также в некоторых случаях усложнение тестирования.
*/
type FacadeAPI struct {
	ToolsService    ToolsService
	GroceryService  GroceryService
	ButcheryService ButcheryService
}

func (f *FacadeAPI) countAllGoods() int {
	sum := f.ToolsService.countAllTools() + f.ButcheryService.countAllButchery() + f.GroceryService.countAllGrocery()
	return sum
}

type ToolsService struct{}

func (t *ToolsService) countAllTools() int {
	return 120
}

type GroceryService struct{}

func (g *GroceryService) countAllGrocery() int {
	return 1220
}

type ButcheryService struct{}

func (b *ButcheryService) countAllButchery() int {
	return 110
}

func FacadeMain() {
	facade := FacadeAPI{}
	allGoods := facade.countAllGoods()
	fmt.Println(allGoods)
}
