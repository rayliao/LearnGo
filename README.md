# GO学习

## 函数

### 函数的声明

函数的声明包括函数名，形式参数列表，返回值列表（可省略）以及函数体
```
func name(parameter-list) (return-list) {
    body
}
```

如果一组形参或返回值有相同类型，不必为每个形参都写出参数类型
```
func f(i, j, k int, s, t string)
func f(i int, j int, k int, s string, t string)
```