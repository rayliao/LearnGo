# GO学习

[Go语言圣经](https://shifei.me/gopl-zh/)

[Go语言圣经](https://docs.ruanjiadeng.com/gopl-zh/)

[Go Web 编程](https://astaxie.gitbooks.io/build-web-application-with-golang/content/zh/index.html)

[Golang学习](http://yougg.github.io/static/gonote/GolangStudy.html)

- [Go基础](#Go基础)
- [结构体](#结构体)
- [函数](#函数)
- [方法](#方法)
- [接口](#接口)
- [并发](#并发)
- [包和工具](#包和工具)

## Go基础

### map

```go
// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
var numbers map[string]int
// 另一种map的声明方式
numbers := make(map[string]int)
numbers["one"] = 1  //赋值
numbers["ten"] = 10 //赋值
numbers["three"] = 3

fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
// 打印出来如:第三个数字是: 3
```
通过`deletes`删除`map`元素
```go
// 初始化一个字典
rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
csharpRating, ok := rating["C#"]
if ok {
    fmt.Println("C# is in the map and its rating is ", csharpRating)
} else {
    fmt.Println("We have no rating associated with C# in the map")
}

delete(rating, "C")  // 删除key为C的元素
```

### make、new操作

make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。

> new返回指针。

> make返回初始化后的（非零）值。

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

### 返回值

如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。这称之为bare return。
```go
func CountWordsAndImages(url string) (words, images int, err error) {
    ...
    return
}
```

### 函数值

函数类型的零值是nil，调用值为nil的函数值会引起panic错误：
```go
var f func(int) int
f(3) // 此处f的值为nil，会引起panic错误
```

函数值可以和nil比较，但函数值之间不可以比较。

## 方法

### 方法声明

在函数声明时，在其名字之前放上一个变量，即是一个方法。

方法可以被声明到任意类型，只要不是一个指针或者一个interface。

## 接口

接口是一组方法的组合。

接口类型是一种抽象的类型。
### interface类型

```go
type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human //匿名字段Human
    school string
    loan float32
}

type Employee struct {
    Human //匿名字段Human
    company string
    money float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
    fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
    fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
    fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
        e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
    s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
    e.money -= amount // More vodka please!!! Get me through the day!
}

// 定义interface
type Men interface {
    SayHi()
    Sing(lyrics string)
    Guzzle(beerStein string)
}

type YoungChap interface {
    SayHi()
    Sing(song string)
    BorrowMoney(amount float32)
}

type ElderlyGent interface {
    SayHi()
    Sing(song string)
    SpendSalary(amount float32)
}
```

interface可以被任意的对象实现。我们看到上面的Men interface被Human、Student和Employee实现。
同理，一个对象可以实现任意多个interface，例如上面的Student实现了Men和YoungChap两个interface。

最后，任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface。
### interface值
```go
package main
import "fmt"

type Human struct {
    name string
    age int
    phone string
}

type Student struct {
    Human //匿名字段
    school string
    loan float32
}

type Employee struct {
    Human //匿名字段
    company string
    money float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
    fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
    fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
        e.company, e.phone)
    }

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
    SayHi()
    Sing(lyrics string)
}

func main() {
    mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
    paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
    sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
    Tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

    //定义Men类型的变量i
    var i Men

    //i能存储Student
    i = mike
    fmt.Println("This is Mike, a Student:")
    i.SayHi()
    i.Sing("November rain")

    //i也能存储Employee
    i = Tom
    fmt.Println("This is Tom, an Employee:")
    i.SayHi()
    i.Sing("Born to be wild")

    //定义了slice Men
    fmt.Println("Let's use a slice of Men and see what happens")
    x := make([]Men, 3)
    //这三个都是不同类型的元素，但是他们实现了interface同一个接口
    x[0], x[1], x[2] = paul, sam, mike

    for _, value := range x{
        value.SayHi()
    }
}
```

你会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，
Go通过interface实现了duck-typing: 即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，
那么这只鸟就可以被称为鸭子"。

### 空interface

所有的类型都实现了空interface。
```go
// 定义a为空接口
var a interface{}
var i int = 5
s := "Hello world"
// a可以存储任意类型的数值
a = i
a = s
```
一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，
如果一个函数返回interface{},那么也就可以返回任意类型的值。

### interface函数参数

interface的变量可以持有任意实现该interface类型的对象，
可以通过定义interface参数，让函数接受各种类型的参数。

举个例子：fmt.Println是我们常用的一个函数，但是你是否注意到它可以接受任意类型的数据。
打开fmt的源码文件，你会看到这样一个定义:
```go
type Stringer interface {
     String() string
}
```
任何实现了String方法的类型都能作为参数被fmt.Println调用:
```go
package main
import (
    "fmt"
    "strconv"
)

type Human struct {
    name string
    age int
    phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
    return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
    Bob := Human{"Bob", 39, "000-7777-XXX"}
    fmt.Println("This Human is : ", Bob)
}
```
### interface变量存储的类型

判断接口变量保存了哪个类型的对象？
Comma-ok断言，判断是否是该类型的变量：
value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
```go
package main

  import (
      "fmt"
      "strconv"
  )

  type Element interface{}
  type List [] Element

  type Person struct {
      name string
      age int
  }

  //打印
  func (p Person) String() string {
      return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
  }

  func main() {
      list := make(List, 3)
      list[0] = 1 //an int
      list[1] = "Hello" //a string
      list[2] = Person{"Dennis", 70}

      for index, element := range list{
          switch value := element.(type) {
              case int:
                  fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
              case string:
                  fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
              case Person:
                  fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
              default:
                  fmt.Println("list[%d] is of a different type", index)
          }
      }
  }
```
### 嵌入interface
```go
type Interface interface {
    sort.Interface //嵌入字段sort.Interface
    Push(x interface{}) //a Push method to push elements into the heap
    Pop() interface{} //a Pop elements that pops elements from the heap
}
```
sort.Interface其实就是嵌入字段，把sort.Interface的所有method给隐式的包含进来了。

另一个例子就是io包下面的 io.ReadWriter ，它包含了io包下面的Reader和Writer两个interface：
```go
// io.ReadWriter
type ReadWriter interface {
    Reader
    Writer
}
```
修改相应的值
```go
var x float64 = 3.4
p := reflect.ValueOf(&x)
v := p.Elem()
v.SetFloat(7.1)
```

## 反射

```go
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
```

## 并发

### goroutine

通过关键字`go`就启动了一个`goroutine`:
```go
go hello(a, b, c)
```
```go
package main

import (
    "fmt"
    "runtime"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        runtime.Gosched()
        fmt.Println(s)
    }
}

func main() {
    go say("world") //开一个新的Goroutines执行
    say("hello") //当前Goroutines执行
}
```

### channels

必须使用`make`创建`channels`
```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```
通过操作符`<-`来接收和发送数据
```go
ch <- v // 发送v到channel ch
v := <- ch // 从ch接收数据并赋值给v
```

## 包和工具

### 包的简介

独立的单元，便于理解和更新，保持独立性，实现封装性。

### 导入路径

每个包是由一个全局唯一的字符串所标识的导入路径定位。出现在import语句中的导入路径也是字符串。
```go
import (
    "fmt"
    "math/rand"
    "encoding/json"

    "golang.org/x/net/html"

    "github.com/go-sql-driver/mysql"
)
```
自定义包的路径从根目录src开始

### 包声明
```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println(rand.Int())
}
```
默认的包名就是包导入路径名的最后一段，因此即使两个包的导入路径不同，它们依然可能有一个相同的包名。
例如，math/rand包和crypto/rand包的包名都是rand。

### 包的匿名导入
```go
import _ "image/png" // register PNG decoder
```

### 工具
src: 用于存储源代码

pkg：用于保存编译后的包的目标文件

bin：用于保存编译后的可执行文件

#### 包文档

包中每个导出的成员和包声明前都应该包含目的和用法说明的注释。
