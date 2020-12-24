package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func appendNilError(f interface{}, err error) (interface{}, error) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		return nil, errors.New("appendNilError: f is not a function")
	}
	param, ret := []reflect.Type{}, []reflect.Type{}
	for i := 0; i < t.NumIn(); i++ {
		param = append(param, t.In(i))
	}
	for i := 0; i < t.NumOut(); i++ {
		ret = append(ret, t.Out(i))
	}
	ret = append(ret, reflect.TypeOf((*error)(nil)).Elem())
	funcType := reflect.FuncOf(param, ret, t.IsVariadic())
	v := reflect.ValueOf(f)
	funcValue := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		results := v.Call(args)
		results = append(results, reflect.ValueOf(&err).Elem())
		return results
	})
	return funcValue.Interface(), nil
}

func main() {
	f := func(a, b int, c string) []string {
		return []string{
			strconv.Itoa(a),
			strconv.Itoa(b),
			c,
		}
	}
	f2, err := appendNilError(f, nil)
	if err != nil {
		fmt.Println(err)
	}

	slist, err := f2.(func(int, int, string) ([]string, error))(5, 3, "hello")
	fmt.Println(slist)
}
