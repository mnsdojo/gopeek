package main

import (
	"fmt"
	"reflect"
)

func Inspect(v any) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	fmt.Println("---Inspect---")
	fmt.Println("Type :", t)
	fmt.Println("Kind:", val.Kind())
	fmt.Println("Value :", val.Interface())

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := t.Field(i)
			value := val.Field(i)
			fmt.Printf("Field %d : %s (%s) = %v\n", i, field.Name, field.Type, value)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			v := val.MapIndex(key)
			fmt.Printf("Key :%v , Value (%s) : %v\n", key, v.Kind(), v.Interface())
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i)
			fmt.Printf("Index %d (%s) : %v\n", i, item.Kind(), item.Interface())
		}
	default:
		fmt.Println("No further inspection available")
	}
}
