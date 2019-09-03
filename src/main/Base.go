package main

import "fmt"

func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ",a) // 取的是值对象
	fmt.Println("&a = ",&a) //值的地址对象
	fmt.Println("*&a = ",*&a) //指向地址的值对象
	fmt.Println("b = ",b)
	fmt.Println("&b = ",&b)
	fmt.Println("*&b = ",*&b)
	fmt.Println("*b = ",*b)
	fmt.Println("c = ",c)
	fmt.Println("*c = ",*c)
	fmt.Println("&c = ",&c)
	fmt.Println("*&c = ",*&c)
	fmt.Println("**c = ",**c)
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)
	fmt.Println("x = ",x)

	// 创建一个map,更像一个字典
	data :=make(map[string]interface{})
	data["userId"]=12
	fmt.Println(data)
}

