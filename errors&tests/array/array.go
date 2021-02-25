package array

type List interface {
	getElement(int) (int, error)
}

type Arr struct {
	values [10]int
}

type OverTheLimitError struct{}
type UnderTheLimitError struct{}

func (OverTheLimitError) Error() string {
	return "the index is over the limit"
}

func (UnderTheLimitError) Error() string {
	return "the index is under the limit"
}

func (a Arr) getElement(n int) (int, error) {

	if n < 0 {
		return 0, UnderTheLimitError{}
	}

	if n > 9 {
		return 0, OverTheLimitError{}
	}

	return a.values[n], nil
}
