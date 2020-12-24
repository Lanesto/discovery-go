package main

import (
	"fmt"
	"reflect"
)

func newMap(key, value interface{}) interface{} {
	keyType := reflect.TypeOf(key)
	valueType := reflect.TypeOf(value)
	mapType := reflect.MapOf(keyType, valueType)
	mapValue := reflect.MakeMap(mapType)
	return mapValue.Interface()
}

func main() {
	m := newMap(" ", 1).(map[string]int)
	m["1"] = 3
	m["5"] = 8
	fmt.Println(m)
}
