package main

import (
	"fmt"
	"net/http"
	"golang_learn/learn"
	"golang_learn/reflect_learn"
)

func sayHello(w http.ResponseWriter,r *http.Request){
	
	fmt.Println(r.Header)
	fmt.Fprint(w,"<h1>你好世界</h1> ")
}




func main() {


	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

	// learn.Test1()
	// learn.Test2()
	// learn.Test3()
	// learn.Test4()
	// learn.Test5()
	// learn.Test8()
	// learn.Test9()
	learn.Test7_6()
	reflect_learn.Test1_1()
	
}
