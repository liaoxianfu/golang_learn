package learn

import (
	"fmt"
)

// 面相对象编程 

// 3.3 匿名组合

/*
go语言的继承是通过组合的文法实现了继承
我们称之为匿名组合
*/

type Base struct {
	Name string
}

func(b *Base) Foo(){
	fmt.Println("foo")
}

func (b *Base)Bar(){
	fmt.Println("Bar")
}


type Foo struct {
	age int
	Base // 在创建实例的时候Base 会自动进行实例化
}


type Bar struct {
	age int 
	*Base //创建实例的时候需要传递Base实例的指针
}




func (f *Foo)Bar(){
	// 重写Bar方法 引用了父类的数据
	f.Base.Bar()
	fmt.Println("Bar Foo")
}


func (b *Bar)foo(){
	b.Base.Foo()
	fmt.Println("数据")
}

func Test6_1(){
	var b = new(Base)
	b.Bar()
	b.Foo()
	var f = new(Foo)
	f.Bar()
	f.Foo()


	var base = new(Base)
	var bar = new(Bar)
	bar.Base = base // 需要显示的传递Base实例的指针
	bar.Foo()
}


// 结构体属性名冲突问题

type X struct {
	Name string
}

type Y struct {
	X
	Name string  // X和Y是不冲突的 因为所有的Y类型能够直接访问的就是Y的最外面一层的Name属性 
}

// 可见性

// 如果需要对外部可见就需要将属性开头字母大写

