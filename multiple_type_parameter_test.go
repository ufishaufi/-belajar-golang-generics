package belajar_golang_generics

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestXxx(t *testing.T) {
	MultipleParameter[string, int]("Shaufi", 100)
	MultipleParameter[int, string](100, "Shaufi")
}
