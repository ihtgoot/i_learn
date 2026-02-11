package helper

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ihtgoot/i_learn/Section_3/internal/config"
)

var app *config.AppConfig

func NewHelper(a *config.AppConfig) {
	a = app
}

// errr 2 kinf : user(clint) and server side

// clint side logging
func ClintError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("clint error !!!!!!!!  STATUS:", status)
	http.Error(w, http.StatusText(status)+"\t funck you", status)
}

// server side logging
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("[ERROR] : \t%s\n\t\t%s", err.Error(), debug.Stack())
	app.ErrorLog.Println("server error !!!!!!\n", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// import (
// 	"errors"
// )

func Getdata() (string, string) {
	o := "Rick Sanches"
	s := "Wubbe Lubba Dup Dup"
	return o, s
}

// divides 2 flat 64 and giver divide by 0 as errro
// func Dividevalue(a, b float64) (float64, error) {
// 	if b == 0 {
// 		err := errors.New("divide by 0")
// 		return 0, err
// 	}
// 	result := a / b
// 	return result, nil
// }

// // adds 2 values int
// func AddValues(a, b int) int {
// 	return a + b
// }

// // // handels divides , takes 2 input and divides them
// func Divide(w http.ResponseWriter, r *http.Request) {
// 	var x, y float64
// 	fmt.Print("x:")
// 	fmt.Scan(&x)
// 	fmt.Print("y:")
// 	fmt.Scan(&y)
// 	f, err := helper.Dividevalue(x, y)
// 	if err != nil {
// 		fmt.Fprintf(w, fmt.Sprintln(err))
// 	} else {
// 		fmt.Fprintf(w, fmt.Sprintf("%f/%f = %f", x, y, f))
// 	}
// }
