package learn

import (
	"fmt"
	"time"
)

func Test8() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04"))
}

// 类型定义和类型别名的区别
/**
自定义类型是定义了一个全新的类型。
我们可以基于内置的基本类型定义，也可以通过struct定义

go 1.9 新增
类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
就像一个孩子小时候有小名、乳名，上学后用学名，
英语老师又会给他起英文名，但这些名字都指的是他本人。


*/
func Test9() {
	type NewInt int  // 类型定义 这时候的类型就是NewInt
	type MyInt = int // 类型别名 类型还是int
	var a MyInt
	var b NewInt
	fmt.Printf("NewInt type is %T\n", b)
	fmt.Printf("MyInt type is %T\n", a)
}

// 结构体

/*
Go语言中的基础数据类型可以表示一些事物的基本属性，
但是当我们想表达一个事物的全部或部分属性时，
这时候再用单一的基本数据类型明显就无法满足需求了，
Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，
这种数据类型叫结构体，英文名称struct。
也就是我们可以通过struct来定义自己的类型了。
Go语言中通过struct来实现面向对象。
*/

type person struct {
	name string
	city string
	age  int
}

// 结构体实例化

/*
只有当结构体实例化时，才会真正地分配内存。
也就是必须实例化后才能使用结构体的字段。
结构体本身也是一种类型，
我们可以像声明内置类型一样使用var关键字声明结构体类型。
*/

func Test10() {
	var p1 person
	p1.name = "liao"
	p1.city = "henan"
	p1.age = 20
	fmt.Println(p1)
	fmt.Printf("%v\n", p1)
}

// 匿名结构体

func Test11() {
	// 定义匿名结构体
	var person struct {
		Name string
		Age  int
	}
	person.Name = "la"
	person.Age = 10
	fmt.Println(person)
}

// 创建指针类型的结构体

func Test12() {
	var p2 = new(person)
	// go中支持直接对结构体指针进行赋值
	p2.age = 10
	p2.city = "he"
	p2.name = "ss"
	fmt.Println(p2)
	fmt.Printf("%T\n", p2)

}

// 取结构体实例化

/*
使用&对结构体进行取地址操作相当于对该结构体类型进行了
一次new实例化操作。

没有初始化的结构体默认是成员变量对应类型的默认值
*/

func Test13() {
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}
}

// 结构体初始化

/**
使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
*/

func demo1() {
	// 直接复赋值
	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
}

func demo2() {
	// 对结构体指针赋值

	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
}

func demo3() {
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
	// 省略键 直接写值
	/*
	   必须初始化结构体的所有字段。
	   初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	   该方式不能和键值初始化方式混用。
	*/
	p8 := &person{
		"沙河娜扎",
		"北京",
		28,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
}

// 构造函数
type animal struct {
	age  int
	name string
}

func newAnimal(age int, name string) *animal {
	return &animal{
		age:  age,
		name: name,
	}
}

func Test14() {
	p1 := newAnimal(10, "dog")
	fmt.Println(*p1)
}

