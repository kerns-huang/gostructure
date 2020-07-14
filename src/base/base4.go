package main

const N = 3

func main() {
	m := make(map[int]*int)
	for i := 0; i < N; i++ {
		m[i] = &i  //代表指针对象，这个指针指向哪里？
	}
	for _, v := range m {
		println(*v) //指针的值对象
	}
}
