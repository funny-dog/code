package main

import (
	"fmt"
)

// 定义数组
var arr [5]int

// 传递函数指针，*号在前
func foo(arrx *[1e6]int) {
	arrx[0] = 5
}

func main() {
	//访问数组元素
	fmt.Printf("%d\n", arr[0])
	//初始化数组
	arr2 := [5]int{3, 4, 5, 6, 7}
	//赋值
	arr2[2] = 35
	fmt.Printf("%d\n", arr2[2])
	//指针数组，并为位置0指向的数据分配空间
	arr3 := [5]*int{0: new(int)}
	*arr3[0] = 20
	fmt.Printf("%d\n", *arr3[0])
	//数组是一个值
	arr = arr2
	fmt.Printf("%d\n", arr[0])
	//复制指针数组，显然是浅复制
	var arr4 [5]*int
	arr4 = arr3
	fmt.Printf("%d %d\n", arr3[0], arr4[0])
	//二维数组
	var arr5 [4][2]int
	arr6 := [4][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	arr5 = arr6
	arr6 = arr5
	//指定某一行
	arr7 := [4][2]int{1: {1, 2}, 3: {3, 4}}
	fmt.Printf("%d\n", arr7[1][1])
	//指定行和列
	arr8 := [4][2]int{0: {0: 6}, 2: {1: 8}}
	fmt.Printf("%d\n", arr8[2][1])
	//复制一行给一维数组
	var arr9 [2]int = arr5[1]
	fmt.Printf("%d\n", arr9)
	//使用指针传递大数组
	var arr10 [1e6]int
	foo(&arr10)
	fmt.Printf("%d\n", arr10[0])
}
