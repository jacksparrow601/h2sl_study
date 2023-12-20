#
##
##### 格式控制
格式化字符串：
- %d: 十进制整数
- %f: 浮点数
- %s: 字符串
- %t: 布尔值
- %v: 自动选择合适的格式
- %T: 输出对应变量的type
```go
package main

import "fmt"

func main() {
    name := "John"
    age := 30
    height := 175.5
    isStudent := false

    fmt.Printf("Name: %s, Age: %d, Height: %.2f, Is Student: %t\n", name, age, height, isStudent)
    // Output: Name: John, Age: 30, Height: 175.50, Is Student: false
}

```
##### defer
在go中，`defer`关键字用于确保函数调用在程序执行结束时执行。通常用于资源清理工作，例如关闭文件、解锁互斥量、发送HTTP响应等。`defer`使代码更加整洁，确保无论函数通过哪条路径退出，都能执行必要的清理工作。

使用场景
- 资源管理：确保（如文件、网路连接等）在不需要的时被释放
- 互斥锁：确保在函数退出时释放锁
- 错误处理： 在函数推出前执行错误检查或清理工作
- 日志记录： 在函数开始和结束时进行日志记录

在函数中使用了`defer`关键字时刻，紧跟其后的函数调用会被推迟执行。无论上一级函数是正常结束，还是因为错误提前退出,这个被推迟的函数调用一定会在当前上一级调用它的函数执行完毕时执行。
```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // 即使在这个位置，即使存在err，函数提前返回，file.Close()函数还是可以正常执行，文件关闭
}
```
1. 当存在多个`defer`语句： 执行顺序会按照先进后出的顺序，即最后一个`defer`先执行，第一个`defer`最后执行
2. 参数的求值： 在`defer`语句中，函数的参数会在`defer`语句被执行时立即求值，而不是在实际调用时求值
3. 与返回值的交互： `defer`语句可以修改函数的返回值，但这需要使用命名的返回值。
