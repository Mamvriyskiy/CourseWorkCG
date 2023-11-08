package menu

import (
	"strconv"
)

type MyError struct {
	err error
	message string
}

func (e *MyError) Error() string {
	return e.message
}

func CheckEntrySize(a, b string) (int, int, MyError, MyError) {
	var errA, errB MyError
	var sizeSceneA, sizeSceneB int

	sizeSceneA, errA.err = strconv.Atoi(a)

	sizeSceneB, errB.err = strconv.Atoi(b)

	if errA.err == nil && errB.err == nil {
		if sizeSceneA <= 0 || sizeSceneA > 100 {
			errA.message = "Введите число из диапазона(от 1 до 100)"
		}

		if sizeSceneB <= 0 || sizeSceneB > 100 {
			errB.message = "Введите число из диапазона(от 1 до 100)"
		}
	} else {
		if errA.err != nil {
			errA.message = "Данные введены неверно"
		}

		if errB.err != nil {
			errB.message = "Данные введены неверно"
		}
	}

	return sizeSceneA, sizeSceneB, errA, errB
}
