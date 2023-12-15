package pattern

import (
	"fmt"
	"sync"
	"time"
)

/* Паттерн Цепочка обязанностей позволяет вынести дорогостоящую бизнес-логику в несколько отдельных этапов, образуя конвейер.
Таким образом, можно добиться ускорения выполнения бизнес-логики.

К плюсам можно отнести относительно несложную реализацию, ускорение работы программы.

К минусам можно отнести то, что иногда построить конвейерную ленту бывает непросто в виду нетривиальности исходного кода.
*/

func ChainMain() {
	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func(ch1 chan int) {
		for i := 0; i < 5; i++ {
			wg.Add(1)
			ch1 <- i
			fmt.Println("Пиво ", i, " отправлено на фильтрацию")
		}
	}(ch1)

	go func(ch1, ch2 chan int) {
		for v := range ch1 {
			time.Sleep(time.Second)
			ch2 <- v
			fmt.Println("Пиво ", v, " отправлено на перемешивание")
		}
	}(ch1, ch2)

	go func(ch2, ch3 chan int) {
		for v := range ch2 {
			time.Sleep(time.Second)
			ch3 <- v
			fmt.Println("Пиво ", v, " отправлено на разлив")
		}
	}(ch2, ch3)

	go func(ch3 chan int) {
		for v := range ch3 {
			time.Sleep(time.Second)
			fmt.Println("Пиво ", v, " готово")
			wg.Done()
		}
	}(ch3)

	wg.Wait()
}
