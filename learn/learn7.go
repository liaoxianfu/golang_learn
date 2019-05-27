package learn

import (
	"fmt"
)

// go语言的接口不同于java 必须强制实现接口中的所有方法
// go语言是非侵入的接口

type File struct {
	Name string
}

type IFile interface {
	Read()
	Write()
	Close()
}

// 接口的实现
func (f *File) Read() {
	fmt.Println(f.Name)
}

func (f *File) Write() {
	fmt.Println(f.Name + "write")
}

func (f *File) Close() {
	fmt.Println(f.Name + "close")
}

func Test7_1() {
	file := new(File)
	file.Name = "file"
	file.Write()
	var ifile IFile
	ifile = file
	ifile.Close()

}

type zhifu interface {
	zhifu(money float64) bool
}

type TaoBao struct {
	Name string
}

type Dangdang struct {
	Name string
}

// 实现接口

// 淘宝的支付接口实现
func (t *TaoBao) zhifu(money float64) bool {
	fmt.Printf("%s 支付了 %.2f元", t.Name, money)

	return true
}

// 当当的支付

func (d *Dangdang) zhifu(money float64) bool {
	fmt.Printf("%s 支付了%.2f元", d.Name, money)
	return true
}

func Test7_2() {
	var zhifu zhifu
	t := new(TaoBao)
	t.Name = "淘宝"
	zhifu = t
	var flag = zhifu.zhifu(10000)
	fmt.Println(flag)

	d := new(Dangdang)
	d.Name = "当当"
	zhifu = d
	flag = zhifu.zhifu(20000)
	fmt.Println(flag)

}

// 接口嵌套
type Tuihuo interface {
	tuihuo(name string) bool
}

type Tuikuan interface {
	tuikuan(money float64) bool
}

type Shouhou interface {
	Tuihuo
	Tuikuan
}

// 接口实现
func (c *TaoBao) tuihuo(name string) bool {
	fmt.Printf("%s 退回了%s商品", c.Name, name)
	return true
}

func (s *TaoBao) tuikuan(money float64) bool {
	fmt.Printf("%s 退钱%.2f", s.Name, money)
	return true
}

func Test7_3() {
	taobao := new(TaoBao)
	taobao.Name = "淘宝"
	var tk Tuikuan
	tk = taobao
	fmt.Println(tk.tuikuan(56))

}

// 空接口


// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
// 空接口类型的变量可以存储任意类型的变量。
func Test7_4() {
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x) //type:string value:Hello 沙河
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x) // type:int value:100
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x) //type:bool value:true

}


// 空接口作为map的值
// 使用空接口实现可以保存任意值的字典。
func Test7_5(){
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

// 类型断言
func Test7_6(){
	var x interface{}
	x="hehe"
	v,ok :=x.(string)
	if ok{
		fmt.Println(v)
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}


}
