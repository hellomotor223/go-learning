package main

import (
	"fmt"
	"time"
)

func main() {
	a:=3
	//普通if
	if a<0 {
		fmt.Print("a时小于0的！")
	}else {
		fmt.Print("a是大于0的！")
	}
	fmt.Println("-----------------------")
	//嵌套if
	if a>1 {
		fmt.Println("a是大于1的！")
		if a<4 {
			fmt.Println("a是小于4的！")
		}
	}

	fmt.Println("-----------siwtch语句------------")
	//switch 语句
	switch 3 {
	case 1:
		fmt.Println("判断为1")
	case 2:
		fmt.Println("判断为2")
	default:
		fmt.Println("不为1也不为2")
	}

	var b interface{}
	b=32
	switch b.(type) {
	case int:
		fmt.Println("类型为整型")
	case string:
		fmt.Println("类型为字符串")
	default:
		fmt.Println("非整型也非字符串")

	}

	fmt.Println("------------for语句--------------")

	//死循环
	//for  {
	//	fmt.Println("我是张小波")
	//	time.Sleep(1*time.Second)
	//}
	//
	//for i:=1;i<=10 ;i++  {
	//	fmt.Println("我是张小波",i)
	//	time.Sleep(1*time.Second)
	//}
	//
	//c:=[]string{"香蕉","苹果","雪梨"}
	//for key,value:=range c  {
	//	fmt.Println("key：",key)
	//	fmt.Println("value：",value)
	//}
	//for _,value:=range c  {
	//	fmt.Println("value：",value)
	//}

	fmt.Println("-------goto------------")
	goto One
	fmt.Println("执行流程")
	One:
		fmt.Println("这是代码块一")

	//死循环
	//Two:
	//	fmt.Println("这是代码块一")
	//	time.Sleep(1*time.Second)
	//goto Two

	fmt.Println("---------break--------------")
	for  {
		fmt.Println("break关键字练习")
		time.Sleep(1*time.Second)
		break
	}

	for i:=1;i<=3 ;i++  {
		for j:=1;j<=2 ;j++  {
			fmt.Println("break2 练习")
			time.Sleep(1*time.Second)
			break
		}
	}

	fmt.Println("---------continue--------------")
	for i:=1;i<=3 ;i++  {
		if i>=2 {
			fmt.Println("当前continue循环体仍然执行")
			continue
		}
		fmt.Println("continue 练习")
	}
}
