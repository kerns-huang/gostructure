package main
//结构体
type S struct {

}
//接口
type IF interface {
	F()
}

// 方法
func (s S) F() {

}

// 返回对象的方法
func initType() S{
	var s S
	return s
}
//返回对象的指针
func initPoint() *S{
	var s *S
	return s
}
// 返回一个 类似 java的 object ，结果不是一样？
func initEFaceType() interface{}{
	var s S
	return s
}
// 返回一个 指针，
func initEFacePointer() interface{}{
	var s *S
	return s
}

// 返回特定的接口
func initIFaceType() IF{
	var s S
	return s
}


func initIFacePointer() IF{
	var s *S
	return s
}

func main()  {
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


