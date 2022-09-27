Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Вывод программы:
<nil>
false

функция Foo возвращает указатель на интерфейс типа os.PathError, в котором лежит значени nil,
Поэтому при сравнении err и nil в функции main мы получаем false (указатель на значение != значение).
```