/*
Пакет реализует вывод общей информации обо всех тренировках и прогулках.
В пакете есть интерфейс, который содержит два метода: Parse() и ActionInfo(),
и функция, в которую будут передаваться слайсы со строками данных и экземпляры структур Training и DaySteps
*/
package actioninfo

import (
	"fmt"
	"log"
)

/*
Интерфейс DataParser, в котором содержатся сигнатуры методов Parse() и ActionInfo() из пакетов trainings и daysteps.
Используется в качестве параметра в функции Info(), чтобы была возможность передавать в неё в качестве аргументов экземпляры структур Training и DaySteps.
*/
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

/*
Функция принимает слайс строк с данными о тренировках или прогулках и экземпляр одной из структур Training или DaySteps
и выводит строку с информацией об активности с помощью метода ActionInfo().
*/
func Info(dataset []string, dp DataParser) {
	for _, x := range dataset {
		err := dp.Parse(x)
		if err != nil {
			log.Println(err)
			continue
		}
		y, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(y)
	}
}
