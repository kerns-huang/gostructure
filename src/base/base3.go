package main

type S struct {

}

func f(x interface{}){

}

func g(x *interface{}){

}

func main()  {
	s := S{}
	p := &s
	f(s)
	//g(s)
	f(p) //指针也是interface
	//g(p) 指针不能代表 指针对象？

}