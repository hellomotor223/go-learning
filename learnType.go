package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//var i uint8 = 1   //1
	//var i byte =1     //1
	//var i int32 = 1   //4
	//var i int = 1      //8  根据操作系统位数回改变
	//var i uint = 1     //8
	//var i float32 = 1   //4
	var i float64 = 1    //8
	fmt.Print(unsafe.Sizeof(i))
}
