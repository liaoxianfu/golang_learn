package reflect_learn

import (
	"fmt"
	"reflect"
)

// 反射是指在程序运行期对程序本身进行访问和修改的能力。
// 程序在编译时，变量被转换为内存地址，
// 变量名不会被编译器写入到可执行部分。
// 在运行程序时，程序无法获取自身的信息。
// 支持反射的语言可以在程序编译期将变量的反射信息，
// 如字段名称、类型信息、结构体信息等整合到可执行文件中，
// 并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。
// Go程序在运行期使用reflect包访问程序的反射信息。


/*
在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组
成的(我们在上一篇接口的博客中有介绍相关概念)。 
在Go语言中反射的相关功能由内置的reflect包提供，
任意接口值在反射中都可以理解为由reflect.Type和reflect.Value
两部分组成，并且reflect包提供了reflect.TypeOf和reflect.ValueOf
两个函数来获取任意对象的Value和Type。
*/



func reflectType(x interface{}){
	t:=reflect.TypeOf(x)
	fmt.Println(t)
	v:=reflect.ValueOf(x)
	fmt.Println(v)
}

func Test1_1(){
	var x int = 1
	reflectType(x)
}


// type name 和type kind

/*
使用type关键字构造很多自定义类型，而种类（Kind）就是指底层的类型，
但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（
Kind）。 举个例子，我们定义了两个指针类型和两个结构体类型，
通过反射查看它们的类型和种类。
*/

type myInt int64

func Test1_2(){
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32
	


}