package main

import "fmt"

func main() {
	a:= byte(0)
	b:= byte(1)
	fmt.Print("a的值位：")
	fmt.Print(a)
	fmt.Print("\n")
	fmt.Print("b的值位：")
	fmt.Print(b)
	fmt.Print("\n")
	fmt.Print("a&b的值位：")
	fmt.Print(a&b)
	fmt.Print("\n")
	fmt.Print("b&b的值位：")
	fmt.Print(b&b)
	fmt.Print("\n")
	fmt.Print("a|b的值位：")
	fmt.Print(a|b)
	fmt.Print("\n")
	fmt.Print("b|b的值位：")
	fmt.Print(b|b)
	fmt.Print("\n")
	fmt.Print("a^b的值位：")
	fmt.Print(a^b)
	fmt.Print("\n")
	fmt.Print("a^a的值位：")
	fmt.Print(a^a)
	fmt.Print("\n")
	fmt.Print("b^b的值位：")
	fmt.Print(b^b)
	fmt.Print("\n")
	fmt.Print("b<< 左移一位 的值位：")
	fmt.Print(b<<1)
	fmt.Print("\n")
	fmt.Print("b>> 右移一位 值位：")
	fmt.Print(b>>1)
	fmt.Print("\n")





}
