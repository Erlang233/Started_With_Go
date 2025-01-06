package main

import (
	"fmt"
	"os"
)

type myObject struct{
	x int
	y int
}

func main(){
	var str string
	fmt.Println(str) // ""

	//一个声明语句中同时声明一组变量
	var i, j, k int
	fmt.Println(j,i,k)
	var b,f,s = true, 2.3, "4"
	fmt.Println(b,f,s)
	//如果省略每个变量的类型，将可以声明多个类型不同的变量

	var function , err = os.Open("") // os.Open returns a file and an error
	fmt.Println(function,err)

	// 指针
	/*
	一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。
	如果用“var x int”声明语句声明一个x变量，那么&x表达式（取x变量的内存地址）将产生一个指向该整数变量的指针，指针对应的数据类型是*int，指针被称之为“指向int类型的指针”。
	如果指针名字为p，那么可以说“p指针指向变量x”，或者说“p指针保存了x变量的内存地址”。
	同时*p表达式对应p指针指向的变量的值。一般*p表达式读取指针指向的变量的值，这里为int类型的值，同时因为*p对应一个变量，所以该表达式也可以出现在赋值语句的左边，表示更新指针所指向的变量的值。
	*/

	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
	erlang := myObject{1,4}
	p := &erlang
	fmt.Println(p,*p, p.x)
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)


	// new 函数

	newP := new(int)
	fmt.Println(newP)
	fmt.Println(*newP)
	*newP = 2
	fmt.Println(newP)
	fmt.Println(*newP)
}
func incr(p *int) int {
    *p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
    return *p
}
func Erlang(x int) int{
	age := x + 3 //在函数内部，有一种称为简短变量声明语句的形式,它以“名字 := 表达式”形式声明变量，变量的类型根据表达式来自动推导 
	//因为简洁和灵活的特点，简短变量声明被广泛用于大部分的局部变量的声明和初始化。
	//简短变量声明左边的变量可能并不是全部都是刚刚声明的。如果有一些已经在相同的词法域声明过了（§2.7），那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了
	age, nxtyear := x + 3, x+4
	fmt.Println(nxtyear)
	//age,nxtyear := 3, 4 // no new variables on left side of :=
	return age
}