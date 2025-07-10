package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

type InsertQuery struct {
	Query string
	Args  []any
}

func InsertStruct(tableNameOverride string, v any) (*InsertQuery, error) {

	val := reflect.ValueOf(v)

	// If an pointer is passed in v then get struct using .Elem()
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("InsertStruct: expected struct, got %s", val.Kind())
	}

	valType := val.Type()

	// Use tableNameOverride as tableName is given
	tableName := tableNameOverride
	if tableName == "" {
		if tag := valType.Field(0).Tag.Get("table"); tag != "" {
			tableName = tag
		} else {
			tableName = SnakeCase(valType.Name())
		}
	}

	var cols, ph []string
	var args []any

	collect := func(field reflect.StructField, fv reflect.Value) {
		dbTag := field.Tag.Get("db")
		switch dbTag {
		case "-":
			return
		case "":
			dbTag = SnakeCase(field.Name)
		}

		if strings.Contains(dbTag, "omitempty") && IsZero(fv) {
			return
		}

		tagParts := strings.Split(dbTag, ",")

		colName := tagParts[0]

		cols = append(cols, colName)
		ph = append(ph, "?")

		args = append(args, toNull(fv).Interface())

	}

	var walk func(reflect.Type, reflect.Value)

	walk = func(t reflect.Type, v reflect.Value) {
		for i := range t.NumField() {
			sf := t.Field(i)
			fv := v.Field(i)

			if sf.Anonymous && fv.Kind() == reflect.Struct {
				walk(fv.Type(), fv)
				continue
			}

			if !sf.IsExported() {
				continue
			}

			collect(sf, fv)

		}
	}

	walk(valType, val)

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(cols, ", "),
		strings.Join(ph, ", "),
	)

	return &InsertQuery{Query: query, Args: args}, nil

}

func SnakeCase(value string) string {
	var b strings.Builder
	for i, r := range value {
		if unicode.IsUpper(r) && i != 0 {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

func IsZero(fv reflect.Value) bool {
	return fv.IsZero()
}

func toNull(v reflect.Value) reflect.Value {

	switch v.Kind() {
	case reflect.String:
		var ns sql.NullString
		if !IsZero(v) {
			ns.Valid = true
			ns.String = v.String()
		}
		return reflect.ValueOf(ns)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var ni sql.NullInt64
		if !v.IsZero() {
			ni.Valid = true
			ni.Int64 = v.Int()
		}
		return reflect.ValueOf(ni)
	case reflect.Float32, reflect.Float64:
		var nf sql.NullFloat64
		if !v.IsZero() {
			nf.Valid = true
			nf.Float64 = v.Float()
		}
		return reflect.ValueOf(nf)
	case reflect.Bool:
		var nb sql.NullBool
		if !v.IsZero() {
			nb.Valid = true
			nb.Bool = v.Bool()
		}
		return reflect.ValueOf(nb)
	default:
		return v
	}

}
