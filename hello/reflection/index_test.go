// 编写函数 walk(x interface{}, fn func(string))，参数为结构体 x，并对 x 中的所有字符串字段调用 fn 函数。难度级别：递归。
package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{{
		"Struct with one string field",
		struct{ Name string }{"Chris"},
		[]string{"Chris"},
	}, {
		"Struct with two string fields",
		struct {
			Name string
			City string
		}{"Chris", "GZ"},
		[]string{"Chris", "GZ"},
	}, {
		"Nested fields",
		Person{"Chris", Profile{33, "GZ"}},
		[]string{"Chris", "GZ"},
	}, {
		"Pointers to things",
		&Person{
			"Chris",
			Profile{33, "London"},
		},
		[]string{"Chris", "London"},
	}, {
		"Slices",
		[]Profile{
			{33, "London"},
			{18, "BeiJing"},
		},
		[]string{"London", "BeiJing"},
	}}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
