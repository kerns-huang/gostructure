package tree

import (
	"fmt"
	"testing"
)

func TestAcAuto(t *testing.T) {
	// 建树
	BuildTree([]string{"ba", "real", "中国"})

	// 设置fail指针
	SetNodeFailPoint()

	// 查找
	fmt.Println(AcAutoMatch("ea"))
	fmt.Println(AcAutoMatch("real"))
	fmt.Println(AcAutoMatch("eal中国"))
	fmt.Println(AcAutoMatch("reaL"))

}
