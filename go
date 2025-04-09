像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
改变i不影响j

当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。 
如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。 

package main
import "fmt"

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
    g, h := 123, "hello"
    fmt.Println(x, y, a, b, c, d, e, f, g, h)
}

const a, b, c = 1, false, "str" //多重赋值
常量还可以用作枚举：

const (
    Unknown = 0
    Female = 1
    Male = 2
)
iota 特殊常量，可以认为是一个可以被编译器修改的常量
第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：

const (
    a = iota
    b
    c
)

位运算符
位运算符对整数在内存中的二进制位进行操作
假定 A = 60; B = 13; 其二进制数转换为：

A = 0011 1100

B = 0000 1101
-----------------
A&B = 0000 1100
A|B = 0011 1101
A^B = 0011 0001

  for true  {
        fmt.Printf("这是无限循环。\n");
    }

Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：
实例
package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
   /* 声明局部变量 */
   var g int = 10

   fmt.Printf ("结果： g = %d\n",  g)
}

Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
import "fmt"
func main() {
   var a int = 10  
   fmt.Printf("变量的地址: %x\n", &a  )
}


指针的基本用法
指针类型是一种特殊的变量，用于存储其他变量的内存地址‌。指针类型是通过在基本数据类型前加上星号（*）来定义的，例如*int表示指向整型值的指针。指针的主要作用是提供对内存地址的直接访问，从而实现对原始数据的操作和修改，这在系统编程和性能优化中非常有用‌
    ‌声明指针变量‌：在Go中，可以通过在类型前加上星号来声明一个指针变量。例如，var ptr *int声明了一个指向整型的指针变量ptr。
    ‌初始化指针‌：通过取变量的地址来初始化指针。例如，ptr = &a将指针ptr初始化为变量a的地址。
    ‌通过指针访问值‌：通过在指针前加星号来访问它所指向的值。例如，*ptr表示访问指针ptr所指向的值。

unsafe.Pointer的类型和用途

‌unsafe.Pointer是Go语言中的一个特殊类型，用于在不同类型的指针之间进行转换‌。它不能进行指针运算，但可以将任意类型的指针转换为unsafe.Pointer，然后再转换回其他类型的指针。这种类型主要用于底层操作和性能优化，因为它可以绕过Go的垃圾回收机制，直接操作内存地址‌3。
import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
}

以上实例执行输出结果为：

a 变量的地址是: 20818a220
ip 变量储存的指针地址: 20818a220
*ip 变量的值: 20

当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

一个指针变量通常缩写为 ptr。



var numbers [5]int
还可以使用初始化列表来初始化数组的元素：
var numbers = [5]int{1, 2, 3, 4, 5}
numbers := [5]int{1, 2, 3, 4, 5}
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：

var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
或
balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
如果设置了数组的长度，我们还可以通过指定下标来初始化元素：

//  将索引为 1 和 3 的元素初始化
balance := [5]float32{1:2.0,3:7.0}
访问数组元素

数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：

var salary float32 = balance[9]
