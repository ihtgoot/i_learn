package helper

import "errors"

func Getdata() (string, string) {
	o := "Rick Sanches"
	s := "Wubbe Lubba Dup Dup"
	return o, s
}

// divides 2 flat 64 and giver divide by 0 as errro
func Dividevalue(a, b float64) (float64, error) {
	if b == 0 {
		err := errors.New("divide by 0")
		return 0, err
	}
	result := a / b
	return result, nil
}
