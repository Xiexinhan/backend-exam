package main

import (
	"fmt"
	"reflect"
)

func swap[T any](a, b T) {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// if a or b is not a pointer, panic , because we need to modify the value at the address
	if va.Kind() != reflect.Ptr || vb.Kind() != reflect.Ptr {
		panic("swap function requires pointer arguments")
	}

	// if a or b is nil, panic , because we cannot dereference a nil pointer
	if va.IsNil() || vb.IsNil() {
		panic("swap function requires non-nil pointer arguments")
	}

	ea := va.Elem()
	eb := vb.Elem()

	// if a or b is not settable, panic , because we need to modify the value at the address
	if !ea.CanSet() || !eb.CanSet() {
		panic("swap function requires settable pointer arguments")
	}

	temp := reflect.New(ea.Type()).Elem()
	temp.Set(ea)
	ea.Set(eb)
	eb.Set(temp)
}

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)

	swap(&a, &b)

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)
}
