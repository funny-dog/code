package main

import (
	"fmt"
)

func foo(slice []int) {
	fmt.Printf("%d\n", slice[5])
}

func main() {
	//如果是0，就是空切片nil
	slice1 := make([]string, 5)
	//表示长度和容量
	// slice2 := make([]string, 5, 7)
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%d\n", slice3[1])
	//不指定值是切片，否则是数组
	slice4 := []string{99: "wlkfh"}
	fmt.Printf("%s\n", slice4[4])
	//赋值
	slice3[1] = 25
	fmt.Printf("%d\n", slice3[1])
	//从1切片到3，长度是3-1=2，容量是5-1=4
	slice3_ := slice3[1:3]
	//新切片其实也是个指针，修改会影响旧数据
	slice3_[0] = 35
	fmt.Printf("%d\n", slice3[1])
	//在容量范围内，增加长度。如果超过容量范围，则复制到另一块内存，并按规则增加容量
	slice3_ = append(slice3_, 60)
	fmt.Printf("%d\n", slice3[3])
	//第三个参数代表限制容量是4-2=2，只能比原有容量小
	slice1_ := slice1[2:3:4]
	fmt.Printf("%s\n", slice1_[0])
	//设置长度和容量一样，这样保证在append时，使用新的内存
	// slice1_ := slice1[2:3:3]
	//合并切片
	slice5 := []int{1, 2}
	slice6 := []int{3, 4}
	fmt.Printf("%v\n", append(slice5, slice6...))
	//迭代切片，其中索引和值是副本
	for index, value := range slice3 {
		fmt.Printf("Index: %d Value: %d ValueAddr: %X ElemAddr: %X\n", index, value, &value, &slice3[index])
	}
	for index := 2; index < len(slice3); index++ {
		fmt.Printf("Value: %d\n", slice3[index])
	}
	//二维切片
	slice7 := [][]int{{0}, {1, 2}}
	slice7[0] = append(slice7[0], 20)
	fmt.Printf("%v\n", slice7[0])
	//函数传参，只有24字节的信息
	slice8 := make([]int, 1e6)
	foo(slice8)
}
