package main

import (
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Error pada field '%s': %s", e.Field, e.Message)
}

func validateUser(name string) error {
	if len(name) == 0 {
		return &ValidationError{Field: "Name", Message: "Nama tidak boleh kosong"}
	}
	return nil
}

func main() {
	err := validateUser("")
	if err != nil {
		fmt.Println("Terjadi error:", err)
	} else {
		fmt.Println("Validasi sukses!")
	}
}
