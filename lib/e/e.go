package e

//тут обработка ошибок

import "fmt"

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

//если вернется не нулевая ошибка.
func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}

	return Wrap(msg, err)
}
