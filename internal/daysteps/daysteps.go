// В этом пакете реализуется функционал для парсинга строки с данными о прогулках и формирования строки с информацией о них. Содержит структуру DaySteps и два экспортируемых метода для неё.
package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	// Формируем константу под итоговую строку
	resultTemplate = "Количество шагов: %v.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n"
)

// Структура содержит все необходимые данные о дневных прогулках: количество шагов, длительность, а также данные из структуры personaldata.Personal, то есть имя, вес и рост пользователя.
type DaySteps struct {
	Steps                 int           // количество шагов
	Duration              time.Duration // длительность прогулки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}

// Метод парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	// Проверяем, что бы полученная строка не была пустой.
	if datastring == "" {
		return errors.New("error in `Parse` function: the resulting string is empty")
	}

	// Разделяем полученную строку на слайс строк и проверяем, чтобы длина слайса была равна 2
	parseStrings := strings.Split(datastring, ",")
	if len(parseStrings) != 2 {
		return errors.New("error in `Parse` function: incorrect number of values received")
	}

	/* Преобразуем первый элемент слайса (количество шагов) в тип int = получаем кол-во шагов из исходной строки.
	И сохраняем значение типа тренировки в соответствующем поле структуры DaySteps - "Steps".*/
	steps, err := strconv.Atoi(parseStrings[0])
	if err != nil {
		return fmt.Errorf("error in `Parse` function: error converting the number of steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("error in `Parse` function: the number of steps must be greater than 0")
	}
	ds.Steps = steps

	/* Преобразуем второй элемент слайса в time.Duration = получаем длительность прогулки из исходной строки.
	И сохраняем значение типа тренировки в соответствующем поле структуры DaySteps - "Duration".*/
	duration, err := time.ParseDuration(parseStrings[1])
	if err != nil {
		return fmt.Errorf("error in `Parse` function: error parsing the duration of the walk: %w", err)
	}
	if duration <= 0 {
		return errors.New("error in `Parse` function: the walking time must be greater than 0")
	}
	ds.Duration = duration

	// Возвращаем отсутствие ошибок, если парсинг прошел успешно.
	return nil
}

// Метод формирует и возвращает строку с данными о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	// Валидируем входные параметры.
	if ds.Personal.Weight <= 0 {
		return "", errors.New("error in `ActionInfo` function: invalid user parameters. Weight must be positive numbers greater than zero")
	} else if ds.Personal.Height <= 0 {
		return "", errors.New("error in `ActionInfo` function: invalid user parameters. Height must be positive numbers greater than zero")
	}

	// Вычисляем дистанцию в метрах и переводим её в километры.
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	// Вычисляем количество калорий, потраченных на прогулке.
	spentСalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Формируем итоговую строку
	finalLine := fmt.Sprintf(resultTemplate, ds.Steps, distance, spentСalories)

	// Возвращаем итоговую строку и возращаем отсутствие ошибок.
	return finalLine, nil
}
