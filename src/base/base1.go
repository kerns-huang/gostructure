package main

import "fmt"

func main()  {
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
