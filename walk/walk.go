package main

import "reflect"

func Walk(obj interface{}, callback func(string)) {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			callback(f.String())
		case reflect.Struct:
			Walk(f.Interface(), callback)
		default:
			continue
		}
	}
}
