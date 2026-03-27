package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	if datastring == "" || len(datastring) <= 0 {
		return errors.New("Datastring is empty")
	}
	arr := strings.Split(datastring, ",")
	if len(arr) < 2 {
		return errors.New("Amount of arguments less 3")
	}
	if len(arr) > 2 {
		return errors.New("Amount of arguments more 3")
	}
	steps, err := strconv.Atoi(arr[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("Steps' amount less or equal zero")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(arr[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("Duration less or equal zero")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories), nil
}
