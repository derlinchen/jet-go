package utils

import (
	"reflect"
)

func Copy(target interface{}, source interface{}) {
	targetVal := reflect.ValueOf(target).Elem()
	sourceVal := reflect.ValueOf(source).Elem()
	sourceType := sourceVal.Type()
	for i := 0; i < sourceVal.NumField(); i++ {
		name := sourceType.Field(i).Name
		if ok := targetVal.FieldByName(name).IsValid(); ok {
			targetVal.FieldByName(name).Set(reflect.ValueOf(sourceVal.Field(i).Interface()))
		}
	}
}
