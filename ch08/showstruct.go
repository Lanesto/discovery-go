package main

import (
	"errors"
	"fmt"
	"reflect"
)

func fieldNames(s interface{}) ([]string, error) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("fieldNames: s is not a struct")
	}
	names := []string{}
	n := t.NumField()
	for i := 0; i < n; i++ {
		names = append(names, t.Field(i).Name)
	}
	return names, nil
}

func main() {
	s := struct {
		A    string
		BC   int
		DDEF struct{}
	}{
		"a", 1, struct{}{},
	}
	fmt.Println(fieldNames(s))
}
