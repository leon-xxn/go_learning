package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 1)
	ch <- struct{}{} //有1个缓冲可以用，无需阻塞，可以立即执行
	go func() {      //子协程1
		time.Sleep(5 * time.Second) //sleep一个很长的时间
		<-ch
	}()

	ch <- struct{}{} //由于子协程1已经启动，寄希望于子协程1帮自己解除阻塞，所以会一直等子协程1执行结束。如果子协程1执行结束后没帮自己解除阻塞，则希望完全破灭，报出deadlock

	go func() { //子协程2
		time.Sleep(5 * time.Second)
		<-ch
	}()
	name := "飞雪无情"

	nameP := &name //取地址

	fmt.Println("name变量的值为:", name)

	fmt.Println("name变量的内存地址为:", nameP)
	coms := buy(100) //采购100套配件

	//三班人同时组装100部手机

	phones1 := build(coms)

	phones2 := build(coms)

	phones3 := build(coms)

	//汇聚三个channel成一个

	phones := merge(phones1, phones2, phones3)

	packs := pack(phones) //打包它们以便售卖

	//输出测试，看看效果

	for p := range packs {

		fmt.Println(p)

	}

	vegetablesCh := washVegetables()
	waterCh := boilWater()
	fmt.Println("已经安排洗菜和烧水了")
	time.Sleep(2 * time.Second)
	fmt.Println("看下菜和水好了")
	vegetables := <-vegetablesCh
	water := <-waterCh
	fmt.Println("准备好了，可以做了：", vegetables, water)
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog(ctx, "【监控狗】")
	}()
	time.Sleep(5 * time.Second)
	stop()
	wg.Wait()
}

func watchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}

// 工序1采购
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装（" + c + "）"
		}
	}()
	return out
}

func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包（" + c + "）"
		}
	}()
	return out
}

func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	wg.Add(len(ins))
	for _, cs := range ins {
		go p(cs)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func washVegetables() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "洗好的菜"
	}()
	return vegetables
}

func boilWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "烧开的水"
	}()
	return water
}

func test() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}
