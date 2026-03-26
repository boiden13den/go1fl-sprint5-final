package trainings

import (
	"errors"
	"fmt"
	"personaldata"
	"spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	if datastring == "" || len(datastring) <= 0 {
		return errors.New("Datastring is empty")
	}
	arr := strings.Split(datastring, ",")
	if len(arr) < 3 {
		return errors.New("Amount of arguments less 3")
	}
	if len(arr) > 3 {
		return errors.New("Amount of arguments more 3")
	}
	steps, err := strconv.Atoi(arr[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("Steps' amount less or equal zero")
	}
	t.Steps = steps
	t.TrainingType = arr[1]
	duration, err := time.ParseDuration(arr[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("Duration less or equal zero")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var caloriesFunc func(int, float64, float64, time.Duration) (float64, error)
	switch t.TrainingType {
	case "Бег":
		caloriesFunc = spentenergy.RunningSpentCalories
	case "Ходьба":
		caloriesFunc = spentenergy.WalkingSpentCalories
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	spentCalories, err := caloriesFunc(t.Steps, t.Weight, t.Height, t.Duration)
	if err != nil {
		return "", fmt.Errorf("calories calculation failed: %w", err)
	}

	return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories), nil
}
