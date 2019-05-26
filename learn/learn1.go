package learn

import (
	"fmt"
)

// defer 语句

/*
Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，
先被defer的语句最后被执行，最后被defer的语句，最先被执行。

在Go语言的函数中return语句在底层并不是原子操作，
它分为给返回值赋值和RET指令两步。
而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
具体如下图所示：

*/

func Test1() {

	fmt.Println("开始")
	defer fmt.Println("1")
	defer fmt.Println("21")
	defer fmt.Println("3")
	fmt.Println("结束")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func Test2() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

// 函数作为变量

func add(x, y int) int {
	return x + y
}

func Test3() {
	f1 := add
	var ret = f1(10, 20)
	fmt.Println(ret)
}

// 函数作为参数传递

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}
func calc(x, y int, op func(int, int) int) int {

	return op(x, y)

}

func Test4() {
	var res = calc(5, 6, sub)
	var res2 = calc(5, 6, add)
	var res3 = calc(5, 6, mul)
	fmt.Printf("减法%d\n加法%d\n乘法%d\n", res, res2, res3)
}

// 匿名函数的使用

func Test5() {
	var add = func(x, y int) (sum int) {
		sum = x + y
		return
	}
	var re = add(5, 6)
	fmt.Println(re)
}

func addper2(x int) func(int) int {
	fmt.Println(x, "aa")
	return func(y int) int {
		fmt.Println(y, "bb")
		return x + y
	}
}

/*
通过传递不同的参数 产生不同的后缀名 在传递文件名获得文件+后缀名
*/
func makeSuffoxFun(suffix string) func(string) string {
	return func(fileName string) string {
		return fileName + suffix
	}
}

// 多函数返回值
func calc2(x int) (func(int) int, func(int) int) {
	add := func(y int) int {
		return x + y
	}
	sub := func(y int) int {
		return x - y
	}
	return add, sub
}

func Test6() {

	// 闭包的计算首先是调用最外层的函数 计算出来得到的结果是一个函数
	// 得到额函数就是具有外部环境数据的内部函数 在进行计算就会得到内部函数的结果
	var d = addper2(10)
	fmt.Println("hhhh")
	fmt.Println(d(20))

	jpgFun := makeSuffoxFun(".jpg")
	fmt.Println(jpgFun("dog"))
	textFun := makeSuffoxFun(".txt")
	fmt.Println(textFun("dog"))

	add, sub := calc2(10)
	fmt.Println(add(20))
	fmt.Println(sub(5))

}

// panic/recover

/*
Go语言中目前（Go1.12）是没有异常机制，
但是使用panic/recover模式来处理错误。 panic可以在任何地方引发，但recover只有在defer调用的函数中有效。 首先来看一个例子：
recover()必须搭配defer使用。
defer一定要在可能引发panic的语句之前定义。
*/


func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func (){
		err:=recover()
		if err!=nil{
			fmt.Println(err)
			fmt.Println("recover in B")
		}
	}()
	var num2 = 0
	num:=1/num2 // 引发异常
	fmt.Println(num)
}

func funcC() {
	fmt.Println("func C")
}
func Test7() {
	funcA()
	funcB()
	funcC()
}