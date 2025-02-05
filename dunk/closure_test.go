package dunk

import (
	"fmt"
	"testing"
	"time"
)

//闭包捕获外部作用域的变量，而不是复制这个变量

/*
应用场景：
1. 闭包可以用来封装变量
2. 闭包可以用来捕获和存储函数的状态
3. 闭包可以作为函数对象或者回调函数

伪全局变量：捕获外部变量，必报实现了全局变量的效果
函数工厂：根据不同的配置参数来动态创建函数
装饰器模式：在不改变原函数的基础上，为函数添加新的功能
回调函数：将函数作为参数传递给其他函数
并发编程：通过闭包在多个goroutine之间共享变量
保存中间态：通过闭包保存函数的中间状态
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestClosure(t *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func square(x int) int {
	return x * x
}

func doubleAndPrint(fn func(int) int) func(int) int {
	return func(x int) int {
		res := fn(x)
		fmt.Println("res:", res)
		return res
	}
}

func TestDecorator(t *testing.T) {
	doubleSquare := doubleAndPrint(square)
	doubleSquare(3)
}

func TaskAsync(input int, callback func(int)) {
	go func() {
		time.Sleep(1 * time.Second)
		result := input * 3
		callback(result)
	}()
}

func TestCallback(t *testing.T) {
	callback := func(result int) {
		fmt.Println("result:", result)
	}
	TaskAsync(2, callback)
	time.Sleep(2 * time.Second)
}

//局部变量分配在堆上还是栈上
//局部变量通常分配在栈上，栈内存用户存储局部变量和函数调用信息。局部变量的生命周期超出函数的作用范围，or局部变量被传递给其他函数可能被持久化，会转移到堆上，逃逸分析
//栈内存的分配和释放是编译器自动管理的，堆内存的分配和释放是手动管理的(malloc new 分配内存)
//函数调用时，局部变量和参数会被压入栈中，函数返回时，这些数据会被弹出栈，栈内存的分配和释放是编译器自动管理的；后进先出的特点，分配速度非常快，大小比较小且固定
//堆内存的分配和释放是手动管理的，需要程序员自己申请和释放内存，分配速度比较慢，大小不固定，可以动态分配内存，但是容易产生内存泄漏和内存溢出
//分配方式：栈是编辑器分配的，堆是手动分配的；局部变量和参数压入栈中；作用范围：栈的内存用于存储局部变量，函数调用参数；堆内存用户存储动态分配的内存快
//存储数据：栈中存储基本数据类型（整数，浮点指针类型变量。结构简单，存储固定大小短期使用的变量；堆通常存储动态分配对象数组或者大块数据

type MyStruct struct {
	field int
}

func (ms MyStruct) ValueReceiverMethod() {
	fmt.Println("ValueReceiverMethod")
}

func (ms *MyStruct) PointerReceiverMethod() {
	fmt.Println("PointerReceiverMethod")
}

//指针接受者的方法可以通过指针调用，也可以通过值调用；指针接收者的方法调用的自动接引用
func TestTPoint(t *testing.T) {
	var s MyStruct
	s.ValueReceiverMethod()
	s.PointerReceiverMethod()
	tmps := &s
	tmps.ValueReceiverMethod()
	tmps.PointerReceiverMethod()
}

//函数中返回局部变量的地址 内存逃逸问题：从栈内存逃逸到堆内存，go的编译器会分析判断局部变量被外部使用，会在堆上分配内存
//指向slice类型对象的指针如果进行了扩容,这个指针可能由于底层数组的原因变得无效
//动态增长：1.18前的版本 len <1024 2倍增长，len>=1024 1.25倍增长；1.18后的版本 len <256 2倍增长，len>=256后 1.25倍+0.75*256 = 192
