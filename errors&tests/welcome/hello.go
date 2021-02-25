package welcome

import "errors"

func sayHello(name string) (string, error) {

	if len(name) == 0 {
		return "", errors.New("Name is empty")
	}

	s := "Hello " + name

	return s, nil
}
