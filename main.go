//程序所属包
package main

//导入依赖包
import "fmt"
//常量
const h2h string  = "黑中黑"
const name  = "我的名字叫：小波"
const(
	cat string = "猫"
	dog = "狗"
)
const apple,banana string  = "苹果","香蕉"
const a1,a2=1,"二"

//变量
var b int
var c int = 456
var(
	d string
	e int
)
func main() {
	b=123
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(h2h)
	fmt.Println(name)
	fmt.Println(cat)
	fmt.Println(dog)

	fmt.Println(apple)
	fmt.Println(banana)

}
