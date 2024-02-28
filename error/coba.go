package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}
	if id != "eko" {
		return &notFoundError{Message: "data not found"}
	}
	return nil
}

// ok
func main() {
	// hasil, err := Pembagian(100, 0)
	// if err == nil {
	// 	fmt.Println("Hasil", hasil)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	err := SaveData("as", nil)
	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error:", validationErr.Error())
		}else if notFoundErr, ok := err.(*notFoundError);ok{
			fmt.Println("not found error:", notFoundErr.Error())
		}else{
			fmt.Println("Unknown Error:", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}

}

// func Pembagian(nilai int, pembagi int) (int, error) {
// 	if pembagi == 0 {
// 		return 0, errors.New("Pembagian dengan NOL")
// 	} else {
// 		return nilai / pembagi, nil
// 	}
// }
