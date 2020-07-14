//你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行

package main

func f1() {
	defer println("f1-begin")
	f2()
	defer println("f1-end")
}
func f2() {
	defer println("f2-begin")
	f3()
	defer println("f2-end")
}
func f3() {
	defer println("f3-begin")
	panic(0) //退出
	defer println("f3-end")
}
func main() {
	f1()
}