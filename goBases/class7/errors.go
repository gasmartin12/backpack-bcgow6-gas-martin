package main

import (
	"errors"
	"fmt"
)

// custom error type
type MyCustomError struct {
	StatusCode int
	Message    string
}

// creation of error. Allways implemet Error() method
func (err *MyCustomError) Error() string {
	return fmt.Sprintf("%s (%d)", err.Message, err.StatusCode)
}

func obtainError(status int) (code int, err error) {
	if status >= 400 {
		return 500,
			&MyCustomError{
				StatusCode: 500,
				Message:    "Boom! </3",
			}
	}
	return 200, nil
}

var errGlobal = errors.New("Global Error")

func getError() error {
	return fmt.Errorf("Data form: %w", errGlobal)
}

type ErrType2 struct{}

func (e ErrType2) Error() string {
	return "Error2"
}
func main() {

	// //otra forma
	// err1 := fmt.Errorf("first error :(")
	// err2 := fmt.Errorf("Second errror: %w", err1)
	// err3 := errors.New("Third error :( : " + err2.Error())

	// fmt.Println(err1)
	// fmt.Println(err2)
	// fmt.Println(err3)

	// //otra forma
	// status, err := obtainError(436)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("Status %d, all ok!", status)

	// // PAQUETE ERRORS
	// // otra forma
	// StatusCode := 300
	// if StatusCode > 399 {
	// 	// la primera en minuscula y sin punto al final
	// 	err := errors.New("boom!")
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("Status %d, all ok!", StatusCode)

	//funciones del paquete errors
	// funcion as()
	err4 := &MyCustomError{
		StatusCode: 502,
		Message:    "Error One",
	}
	err5 := &MyCustomError{
		StatusCode: 406,
		Message:    "Error Two",
	}
	bothErrorEqual := errors.As(err4, &err5)
	fmt.Println(bothErrorEqual)

	//funcion is
	err6 := getError()
	coincidence := errors.Is(err6, errGlobal)
	println(coincidence)

	//Unwrap()
	err7 := ErrType2{}
	err8 := fmt.Errorf("error1: %w", err7)

	fmt.Println(errors.Unwrap(err8))
}
