# GO学习

[Go语言圣经](https://shifei.me/gopl-zh/)

## Table of contents

- [结构体](#结构体)
- [函数](#函数)

## 结构体

### 定义

结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。

### 声明

如下代码，创建了一个自定义类型`person`代表一个实体，此实体拥有属性：姓名和年龄
```go
// * 一个string类型的字段name，用来保存用户名称这个属性
// * 一个int类型的字段age，用来保存年龄这个属性
type person struct {
	name, gender string // 相邻成员类型相同，可以合并到一行
	age  int
}
```

如果成员名字以大写字母开头，就是可导出的，一个结构体可能同时包含导出和未导出的成员

### 使用

```go
var P0 person

// 通过点操作符访问
P0.name = "ray"
P0.age = 18
fmt.Printf("The person's name is %s", P0.name)

// 对成员取地址，通过指针访问
position := &P0.name
*position += "liao"

// 也可以指向实体
lord := &P0
lord.age = 200
```

* 1.按照顺序提供初始值

```go
P1 := person{"cong", 35}
fmt.Printf("\n%s's age is %v", P1.name, P1.age)
```

* 2.通过名值对的方式初始化，可以任意顺序

```go
P2 := person{age: 3, name: "yoyo"}
fmt.Printf("\n%s's age is %d\n", P2.name, P2.age)
```

* 3.通过`new`函数分配一个指针

```go
P3 := new(person)
P3.name = "Danny"
fmt.Println(P3)
```

### 匿名字段

只提供类型，不写字段名的方式，亦称嵌入字段

```go
type Human struct {
    name, phone string
    age, weight int
}

type Skills []string

type Student struct {
    Human // 匿名字段，那么默认Student就包含了Human的所有字段
    speciality string
    phone string
    Skills
    int // 内置类型作为匿名字段
}

// 初始化一个学生
long := Student{Human{"long", "1368888", 60, 80}, "jerk off", []string{"code"}, 88}

fmt.Println("His name is ", long.name)
fmt.Println("His speciality is ", long.speciality)

// 字段的继承
long.weight += 20
fmt.Println("long become more fat")
fmt.Println("His weight is ", long.weight)

// 访问匿名字段
long.Human = Human{"amy", 2, 20}
long.Human.age--

// 进行函数操作
long.Skills = append(long.Skills, "diy", "golang")

// 修改匿名内置类型
long.int = 100

fmt.Println(long)
```

这里有一个问题：如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，
那么该怎么办呢？

### 比较

```go
type Point struct{ x, y int }

p := Point{1, 2}
q := Point{2, 1}
fmt.Println(p.x == q.x && p.y == q.y)
fmt.Println(p == q)
```

### 组合函数

```go
type Rect struct {
    width, length float64
}

func (rect Rect) area() float64 {
    return rect.width * rect.length
}

var rect = Rect{100, 200}

fmt.Println("Width:", rect.width, "Length:", rect.length, "Area:", rect.area())
```

## 函数

### 函数的声明

函数的声明包括函数名，形式参数列表，返回值列表（可省略）以及函数体
```go
func name(parameter-list) (return-list) {
    body
}
```

如果一组形参或返回值有相同类型，不必为每个形参都写出参数类型
```go
func f(i, j, k int, s, t string)
func f(i int, j int, k int, s string, t string)
```
没有函数体的函数声明

```go
func Sin(x float64) float //implemented in assembly language
```
