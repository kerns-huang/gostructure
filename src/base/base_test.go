package main

import (
	"fmt"
	"testing"
)

func TestBase1(t *testing.T){
	s :=[]int{1,2,3}
	ss :=s[1:]
	ss = append(ss,4)
	for _,v := range ss{
		//这种情况下 v 相当于被重新拷贝复制了
		v+=10
	}
	for i := range  ss{
		// 这种情况下，是数组元素的重新定义
		ss[i]+=10
	}
	fmt.Print(ss)
}

func TestBase2(t *testing.T)  {
	//nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.
	//Type must be a pointer, channel, func, interface, map, or slice type
	//println(initType() == nil)
	//下面的结果其实说明，返回结果定义为interface 和 直接定义为特定的对象还是有区别的。
	//虽然概念上该觉是一样
	println(initPoint())
	println(initEFaceType())
	println(initEFacePointer())
	println(initIFaceType())
	println(initIFacePointer())
}

func TestBase3(t *testing.T)  {
	s := S{}
	p := &s
	f(s)
	//g(s)
	f(p) //指针也是interface
	//g(p) 指针不能代表 指针对象？
}