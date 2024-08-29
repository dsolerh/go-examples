package pkg

import "time"

func UseInterface(i Interface) (string, error) {
	str, err := i.GetValue(1523, "string", time.Second)
	return str, err
}
