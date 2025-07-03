package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Расчет калорий при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Валидируем входные параметры.
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid calculation parameters: all values (steps, weight, height, and duration) must be positive numbers")
	}

	// Рассчитываем среднюю скорость с помощью MeanSpeed.
	averageSpeed := MeanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты.
	durationMinutes := duration.Minutes()

	// Рассчитываем и возвращаем количество калорий.
	return ((weight * averageSpeed * durationMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

// Расчет калорий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Валидируем входные параметры.
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("incorrect calculation parameters")
	}

	// Рассчитываем среднюю скорость с помощью MeanSpeed.
	averageSpeed := MeanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты.
	durationMinutes := duration.Minutes()

	// Рассчитываем и возвращаем количество калорий.
	return (weight * averageSpeed * durationMinutes) / minInH, nil
}

// Функция для расчета средней скорости.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Валидируем входные параметры.
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	// Вычисляем дистанцию с помощью distance.
	distance := Distance(steps, height)

	// Преобразуем продолжительность прогулки в часы, вычисляем и возвращаем среднюю скорость.
	return distance / duration.Hours()
}

// Функция для расчета дистанции.
func Distance(steps int, height float64) float64 {
	// Валидируем входные параметры.
	if steps <= 0 || height <= 0 {
		return 0
	}

	// Рассчитываем и возращаем дистанцию в километрах.
	return (height * stepLengthCoefficient * float64(steps)) / mInKm
}
