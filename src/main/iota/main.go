package main

const  a  = iota
const b  = iota

const(
	c=iota
	d=iota
)
//跳值使用法
const(
	e1=iota
	e2=iota
	_
	e3=iota
)
//插队使用法
const(
	f1=iota
	f2=3.14
	f3=iota
)

//表达式隐式使用法;
const(
	g1=iota*2
	g2
	g3
)

//表达式隐式使用法;
const(
	h1=iota*2
	h2=iota*3
	h3
	h4
)
//单行使用法;
const(
	j1,j2=iota,iota+3
	j3,j4
	j5=iota
)




func main() {
	//只能在常量定义中使用
	//fmt.Print(iota) //报错
	println("a的常量值为：",a) //0
	println("b的常量值为：",b) //0

	println("c的常量值为：",c) //0
	println("d的常量值为：",d) //1 const中每新增一行常量声明将使iota计数一次;

	println("e3的常量值为：",e3) //3 跳值使用法

	println("f3的常量值为：",f3) //2 插队使用法

	println("g1的常量值为：",g1) //0 表达式隐式使用法;
	println("g2的常量值为：",g2) //2 表达式隐式使用法;
	println("g3的常量值为：",g3) //4 表达式隐式使用法;

	println("h1的常量值为：",h1) //0 表达式隐式使用法;
	println("h2的常量值为：",h2) //3 表达式隐式使用法;
	println("h3的常量值为：",h3) //6 表达式隐式使用法;
	println("h4的常量值为：",h4) //9 表达式隐式使用法;

	println("j1的常量值为：",j1) //0 单行使用法;
	println("j2的常量值为：",j2) //3 单行使用法;
	println("j3的常量值为：",j3) //1 单行使用法;
	println("j4的常量值为：",j4) //4 单行使用法;
	println("j5的常量值为：",j5) //2 单行使用法;


}
