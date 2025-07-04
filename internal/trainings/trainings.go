/*
В этом пакете реализован функционал по разбору строки с данными о тренировках и формирования строки с информацией о них.
Для этого реализуется два метода: первый для парсинга входящей строки с данными; а второй для формирования информации о тренировках.
*/
package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	// Формируем константу под итоговую строку
	resultTemplate = "Тип тренировки: %v\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"
)

// Структура содержит все необходимые данные о тренировке: количество шагов, тип тренировки, длительность тренировки, а также данные из структуры personaldata.Personal, то есть имя, вес и рост пользователя.
type Training struct {
	Steps                 int           // количество шагов, проделанных за тренировку.
	TrainingType          string        // тип тренировки(бег или ходьба).
	Duration              time.Duration // длительность тренировки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata
}

// Метод парсит строку с данными формата "3456,Ходьба,3h00m" и записывает данные в соответствующие поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {
	// Разделяем полученную строку на слайс строк и проверяем, чтобы длина слайса была равна 3
	parseStrings := strings.Split(datastring, ",")
	if len(parseStrings) != 3 {
		return errors.New("3 parameters are expected: steps, type, and duration")
	}

	/* Преобразуем первый элемент слайса (количество шагов) в тип int = получаем кол-во шагов из исходной строки.
	И сохраняем значение типа тренировки в соответствующем поле структуры Training - "Steps".*/
	steps, err := strconv.Atoi(parseStrings[0])
	if err != nil {
		return fmt.Errorf("incorrect number of steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("the number of steps must be greater than 0")
	}
	t.Steps = steps

	// Сохраняем значение типа тренировки в соответствующем поле структуры Training - "TrainingType".
	t.TrainingType = parseStrings[1]

	/* Преобразуем третий элемент слайса в time.Duration = получаем длительность прогулки из исходной строки.
	И сохраняем значение типа тренировки в соответствующем поле структуры Training - "Duration".*/
	duration, err := time.ParseDuration(parseStrings[2])
	if err != nil {
		return fmt.Errorf("incorrect duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("the duration of the walk must be greater than 0")
	}
	t.Duration = duration

	// Возвращаем отсутствие ошибок, если парсинг прошел успешно.
	return nil
}

// Метод формирует и возвращает строку с данными о тренировке, исходя из того, какой тип тренировки был передан.
func (t Training) ActionInfo() (string, error) {
	// Вычисляем дистанцию и скорость.
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var calories float64
	var err error

	// Проверяем, какой вид тренировки был передан в строке, которую парсили в функции parseTraining. Потом каждого из видов тренировки рассчитываем дистанцию, среднюю скорость и калории.
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error calculating running calories: %w", err)
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error in calculating walking calories: %w", err)
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	// Форматируем результат
	result := fmt.Sprintf(resultTemplate, t.TrainingType, float64(t.Duration.Hours()), distance, meanSpeed, calories)

	// Возвращаем итоговую строку.
	return result, nil
}
