package learn

import (
	"fmt"
)

func Test5_1() {

	// go语言中的数组也是值类型 而c是引用类型

	var arr1 = [...]int{1, 2, 3}
	var arr2 = arr1
	var arr3 = &arr1
	arr1[1] = 55
	fmt.Println(arr1, arr2, *arr3) // [1 55 3] [1 2 3] [1 55 3]
}

// go语言只保留了组合composition特性 放弃了继承等面相对象的特征

// 定义结构体

type Rect struct {
	width, height float64
}

// 定义方法
func (r *Rect) Area() float64 {
	return r.height * r.width
}

// 初始化的方式
func Test5_2(){
	rect1:=new(Rect) // 没有显示声明默认是对应的零值
	fmt.Println(rect1)
	rect2:=&Rect{}
	fmt.Println(rect2)
	rect3:=&Rect{width:10,height:20}
	fmt.Println(rect3.Area())
}




