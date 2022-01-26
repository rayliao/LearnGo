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
	}, {
		"Arrays",
		[2]Profile{
			{22, "London"},
			{11, "GZ"},
		},
		[]string{"London", "GZ"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
		// gotNum := 0
		// for _, x := range got {
		// 	if x == "Bar" || x == "Boz" {
		// 		gotNum++
		// 	}
		// }

		// if gotNum != 2 {
		// 	t.Errorf("expected %d got %d", 2, gotNum)
		// }
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didnt", haystack, needle)
	}
}
