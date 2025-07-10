package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	ID     int    `json:"-"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Age    int    `json:"years_alive_till_now"`
	Income int    `json:"income_amt,omitempty"`
	Address
}

type Address struct {
	Country string `json:"country"`
	city    string `json:"city"`
}

func StructToMap(input any) (map[string]any, error) {

	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("StructToMap: Expected struct received: %s", v.Kind())
	}

	t := v.Type()

	resMap := make(map[string]any)

	walk(v, t, resMap)

	return resMap, nil

}

func collect(fv reflect.Value, field reflect.StructField, resp map[string]any) {

	jsonFieldTag := field.Tag.Get("json")

	if strings.Contains(jsonFieldTag, "-") {
		return
	} else if jsonFieldTag == "" {
		jsonFieldTag = SnakeCase(field.Name)
	}

	if strings.Contains(jsonFieldTag, "omitempty") && IsZero(fv) {
		return
	}

	tagParts := strings.Split(jsonFieldTag, ",")
	if len(tagParts) > 1 {
		jsonFieldTag = tagParts[0]
	}

	resp[jsonFieldTag] = fv.Interface()

}

func walk(v reflect.Value, t reflect.Type, resMap map[string]any) {

	for i := range t.NumField() {
		sf := t.Field(i)
		fv := v.Field(i)

		if !sf.IsExported() {
			continue
		}

		if sf.Anonymous && fv.Kind() == reflect.Struct {
			nestedMap := make(map[string]any)
			walk(fv, fv.Type(), nestedMap)

			fmt.Println(nestedMap)

			resMap[sf.Name] = nestedMap

			continue

		}

		collect(fv, sf, resMap)

	}

}
