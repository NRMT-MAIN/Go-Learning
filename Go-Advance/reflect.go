package main

import (
	"fmt"
	"reflect"
)

func structToMap(s interface{}) map[string]interface{} {
    v := reflect.ValueOf(s)
    t := v.Type()
    
    // Handle pointer
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = v.Type()
    }
    
    if v.Kind() != reflect.Struct {
        return nil
    }
    
    result := make(map[string]interface{})
    
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        
        // Skip unexported fields
        if !field.IsExported() {
            continue
        }
        
        // Use json tag if available
        key := field.Name
        if tag := field.Tag.Get("json"); tag != "" {
            key = tag
        }
        
        result[key] = value.Interface()
    }
    
    return result
}

func reflectExample() {
	x := 42
	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println("Value:", v)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	fmt.Println("Is Zero:" , v.IsZero())
	fmt.Println("Is Struct:", t.Kind() == reflect.Struct)
	fmt.Println("Is Int:", t.Kind() == reflect.Int)
	fmt.Println("Is String:", t.Kind() == reflect.String)

	y := 10
	v2 := reflect.ValueOf(&y).Elem()
	fmt.Println("v2 Value:", v2)
	fmt.Println("v2 Type:", v2.Type())
	fmt.Println("v2 Kind:", v2.Kind())
	fmt.Println("v2 Is Zero:" , v2.IsZero())
	fmt.Println("v2 Is Struct:", v2.Type().Kind() == reflect.Struct)
	fmt.Println("v2 Is Int:", v2.Type().Kind() == reflect.Int)
	fmt.Println("v2 Is String:", v2.Type().Kind() == reflect.String)

	type Person struct {
        Name  string `json:"name"`
        Age   int    `json:"age"`
        Email string `json:"email"`
    }
    
    p := Person{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }
    
    m := structToMap(p)
    fmt.Printf("%+v\n", m)
}