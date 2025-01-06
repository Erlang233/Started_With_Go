// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os" // 包以跨平台的方式，提供了一些与操作系统交互的函数和变量
	"strings"
)

func main() {
    var s, sep string

/*
	s := ""
	var s string
	var s = ""
	var s string = ""
	用哪种不用哪种，为什么呢？
	第一种形式，是一条短变量声明，最简洁，但只能用在函数内部，而不能用于包变量。
	第二种形式依赖于字符串的默认初始化零值机制，被初始化为 ""。
	第三种形式用得很少，除非同时声明多个变量。
	第四种形式显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了。
	实践中一般使用前两种形式中的某个，初始值重要的话就显式地指定变量的值，否则指定类型使用隐式初始化。
*/
    for i := 0; i < len(os.Args); i++ { // := short variable declaration
        s += sep + os.Args[i] // 程序的命令行参数可从 os 包的 Args 变量获取；os 包外部使用 os.Args 访问该变量。
		//os.Args 变量是一个字符串（string）的 切片（slice）
        sep = " "
		//--i --/++ + variable也是不合法的
    }
    fmt.Println(s)

	fmt.Println("第二种模式")

	s, sep = "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = "\n"
    }
    fmt.Println(s)
	/*
	如果在循环中 不需要索引（只需要值），Go 的语法依然会要求你提供一个变量来接收索引。
	nums := []int{10, 20, 30}
	for temp, val := range nums {
		fmt.Println(val) // 忽略 temp，但 temp 未被使用,因此会报错
	}
	

	编译器会报错：temp declared but not used。因为 Go 语言设计上不允许定义未使用的局部变量。
	*/

	fmt.Println("模式3")
	
	fmt.Println(strings.Join(os.Args[1:], "\n"))
}

/*
os.Args 的第一个元素：os.Args[0]，是命令本身的名字；
其它的元素则是程序启动时传给它的参数。s[m:n] 形式的切片表达式，产生从第 m 个元素到第 n-1 个元素的切片，下个例子用到的元素包含在 os.Args[1:len(os.Args)] 切片中。
如果省略切片表达式的 m 或 n，会默认传入 0 或 len(s)，因此前面的切片可以简写成 os.Args[1:]。
*/