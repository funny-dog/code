package main

import (
	"fmt"
)

//映射将键的字符串，通过散列函数转化为散列值
//散列值的高8位用数组存储，用于选择数据存储在哪个桶，桶是哈希表/哈希链表结构
//键值对用字节存储，先存储所有键，再存储所有值
//引用类型不能作为键，比如切片，函数

func main(){
	//使用make定义
	// dict1 := make(map[string]int)
	//初始化
	dict2 := map[string]string{"Orange":"wefi","Red":"wefha"}
	//没有这个键则返回空字符串或对应0值
	fmt.Printf("%s\n",dict2["Blue"])
	dict2["Blue"] = "owuehfauw"
	fmt.Printf("%s\n",dict2["Blue"])
	//删除键
	delete(dict2,"Red")
	//这是同样的效果，相当于传递引用
	removeDict2(dict2, "Red")
	for key, value := range dict2{
		fmt.Printf("Key: %s, Value: %s\n",key,value)
	}
}

func removeDict2(dict map[string]string, key string){
	delete(dict, key)
}