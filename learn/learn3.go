package learn

import (
	"fmt"
)

// 方法和接收者

/*

Go语言中的方法（Method）是一种作用于特定类型变量的函数。
这种特定类型变量叫做接收者（Receiver）。
接收者的概念就类似于其他语言中的this或者 self。
*/

//Person 结构体
type Person struct {
	name string
	age  int
}

//NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// 与普通的函数相比 前面多了（结构体参数）
// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。

func (p Person) Say(info string) string {
	return p.name+info
}
func (p Person) setAge(age int){
	p.age = age
	fmt.Println("之前")
	fmt.Println(p)
}

func Test15(){
	p1:=NewPerson("hh",15)
	fmt.Printf("%T\n",p1)
	var say = p1.Say("我说话了")
	fmt.Println(say)
	p1.setAge(20)
	fmt.Println("之后")
	fmt.Println(p1)// 不会改变结构体的数据
}

// 指针类型的接受者

func (p *Person) setName(name string){
	p.name = name
	fmt.Println("之前")
	fmt.Println(p)
}

func Test16(){
	p1:=NewPerson("hh",15)
	fmt.Println(p1)
	p1.setName("lla")
	fmt.Println("之后")
	fmt.Println(p1)
}

// 自定义构造函数

func newPerson1(name string,age int) *Person{
	p:=new(Person)
	p.age = age
	p.name = name
	return p
}


func Test17(){
	var p = newPerson1("ww",15)
	fmt.Println(p)
}

// 结构体的匿名字段

/*

结构体允许其成员字段在声明时没有字段名而只有类型，
这种没有名字的字段就称为匿名字段。
匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，
因此一个结构体中同种类型的匿名字段只能有一个。
*/

// 嵌套结构体
// 结构体的继承

// 结构体的可见性

// 结构体中字段大写开头表示可公开访问，
// 小写表示私有（仅在定义当前结构体的包中可访问）。



