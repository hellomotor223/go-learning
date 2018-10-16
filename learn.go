//程序所属包
package main

//导入包
import "fmt"
//常量定义
const  NAME string  = "joinu"
//全局变量的声明与赋值
var a string = "加优科技"
//一般类型声明
type joinuInt int
//结构的声明
type Learn struct {
}
//声明接口
type Ilearn interface {
}
//函数声明
func learnJoinu()  {
	fmt.Print("learnJoinu")
}

/*
main func 多行注释
*/
func main() {
	learnJoinu()
	fmt.Println("learn1")
	fmt.Println(a)
	fmt.Println(NAME)
}
