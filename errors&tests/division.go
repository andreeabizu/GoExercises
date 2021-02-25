package division

type DivideByZeroError struct {
}

func (DivideByZeroError) Error() string {

	return "Can't perform division by 0"
}

func Divide(x, y int) (int, error) {

	if y == 0 {
		return 0, DivideByZeroError{}
	}

	return x / y, nil
}
