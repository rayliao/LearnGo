package modals

import "fmt"

// Person 结构体，实体
// 如何声明？
// 声明一个新的类型，作为其他类型的属性或字段的容器
// * 一个string类型的字段name，用来保存用户名称这个属性
// * 一个int类型的字段age，用来保存年龄这个属性
type Person struct {
	name string
	age  int
}

// Human struct
type Human struct {
	name, phone string
	age, weight int
}

// Skills 所有的内置类型和自定义类型都可作为匿名字段
type Skills []string

// Student 包含匿名字段，AKA 嵌入字段
type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
	phone      string
	Skills
	int // 内置类型作为匿名字段
}

func older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

// Log msg
func Log() {
	var P0 Person

	// 通过点操作符访问
	P0.name = "ray"
	P0.age = 18

	position := &P0.name
	*position += "liao"

	lord := &P0
	lord.age = 200

	fmt.Println(*lord)

	// 通过new函数分配一个指针？？？
	P3 := new(Person)
	P3.name = "Danny"
	fmt.Println(P3)
}
