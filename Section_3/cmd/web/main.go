package main

import (
	"fmt"
	"net/http"

	"github.com/ihtgoot/i_learn/Section_3/pkg/handlers"
	"github.com/ihtgoot/i_learn/Section_3/pkg/helper"
)

const portNumber string = ":8080"

// adds 2 values int
func addValues(a, b int) int {
	return a + b
}

// handels divides , takes 2 input and divides them
func Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float64
	fmt.Print("x:")
	fmt.Scan(&x)
	fmt.Print("y:")
	fmt.Scan(&y)
	f, err := helper.Dividevalue(x, y)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintln(err))
	} else {
		fmt.Fprintf(w, fmt.Sprintf("%f/%f = %f", x, y, f))
	}
}

func main() {
	fmt.Println("start server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var owner, saying string
		owner, saying = helper.Getdata()
		n, err := fmt.Fprint(w, fmt.Sprintf("<html>this is the about page <br> %s said \" %s \" %d times</html>", owner, saying, addValues(4949, 5948)))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("\nBytes Written: n : %d", n))
	})

	//http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprint(w, "hi")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Printf(fmt.Sprintf("\nBytes Writen :n : %d", n))
	//})
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/About", handlers.About)
	http.HandleFunc("/Divide", Divide)

	fmt.Println("we started on port number : ", portNumber)
	http.ListenAndServe(portNumber, nil)
}

// Fprintf , write the string in a io buffer ,here the string is "hello" and the io buffer is the http.ResponseWriter i.e w
// on successful execution n stores an int value , stores the no. of bytes written to response writer, and err is left empty awhile if there is an error err store errro message
// in HandelFunc first is the dir of out server that is handles and second arguiment is the fnction that defines how it respaonse if the dir is called ; techinically HandelFunc assign a handler to the dir on access firstis the dir and second is the function that handles the dir
// since go can have more than one return value it provides us with the functionality to handle these during operation without crashig in production this make for a simplifies apprch to error handeling
