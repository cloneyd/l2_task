Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Программа выведет числа от 1 до 8 включительно, и далее будет выводить 0 бесконечно.

В функции merge отсутствует проверка флага ok, который возвращает оператор <-.
Оператор <- всегда возвращает два значения - значение из канала, и флаг, сигнализирующий о состоянии канала.
В случае, если канал будет закрыт первое значение всегда будет равно дефолтному значению типа канала,
и соответсвенно все блоки case всегда будут возвращать true.
```