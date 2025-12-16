package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func TrimAllStrings(a any) {
	if a == nil {
		return
	}
	seen := make(map[uintptr]bool)
	walk(reflect.ValueOf(a), seen)
}

func walk(v reflect.Value, seen map[uintptr]bool) {
	if !v.IsValid() {
		return
	}

	for v.Kind() == reflect.Interface {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return
		}
		p := v.Pointer()
		if p != 0 {
			if _, ok := seen[p]; ok {
				return
			}
			seen[p] = true
		}
		walk(v.Elem(), seen)

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)

			// only string and can set value do trim
			if f.Kind() == reflect.String && f.CanSet() {
				f.SetString(strings.TrimSpace(f.String()))
				continue
			}

			walk(f, seen)
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i), seen)
		}

	case reflect.Map:
		if v.IsNil() {
			return
		}
		for _, k := range v.MapKeys() {
			val := v.MapIndex(k)
			if !val.IsValid() {
				continue
			}
			newVal := reflect.New(val.Type()).Elem()
			newVal.Set(val)

			walk(newVal, seen)

			v.SetMapIndex(k, newVal)
		}
	}
}

func main() {
	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(&a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(&a)

	fmt.Println(a.Next.Next.Name == "name")
}
