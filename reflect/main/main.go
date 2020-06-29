package main

import (
	"fmt"
	"log"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func (p Person) PersonSet(name string, age int, sex string) {
	p.Name = name
	p.Age = age
	p.Sex = sex
	fmt.Println(p)
}
func (p Person) ShowPerson() {
	fmt.Println(p)
}

func (p Person) im() {
	fmt.Println(p)
}

func reflectOfStruct(a interface{}) {
	// 获取reflect.Type类型
	typeObj := reflect.TypeOf(a)
	// 获取reflect.Value类型
	valueObj := reflect.ValueOf(a)
	// 获取Kind类别,下面两种方法都能获取
	KindType := typeObj.Kind()
	//KindValue := valueObj.Kind()
	//fmt.Println(KindType)
	//fmt.Println(KindValue)
	if KindType != reflect.Struct {
		log.Fatal("Kind is not error")
		return
	}
	// 获取字段数量
	fieldsNum := valueObj.NumField()
	for i := 0; i < fieldsNum; i++ {
		fmt.Printf("field %d %v\n", i, valueObj.Field(i))
		// 获取指定的标签值
		tagValue := typeObj.Field(i).Tag.Get("json")
		if tagValue != "" {
			fmt.Printf("field %d tag = %v\n", i, tagValue)
		}
	}
	// 获取方法数量
	MethodNum := valueObj.NumMethod()
	fmt.Printf("has %d methods\n", MethodNum)
	// 调用第一个方法
	valueObj.Method(1).Call(nil)
	// 对有参数的方法调用
	var params []reflect.Value
	params = append(params, reflect.ValueOf("lisi"))
	params = append(params, reflect.ValueOf(88))
	params = append(params, reflect.ValueOf("man"))
	// 传递参数,调用指定方法名的方法
	valueObj.MethodByName("PersonSet").Call(params)
}
func main() {
	var a Person = Person{
		Name: "zhangsan",
		Age:  99,
	}
	// 调用函数
	reflectOfStruct(a)
}
