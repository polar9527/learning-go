package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func PrintUtil(message string) {
	fmt.Println(message, " : ", time.Now().Format("2006-01-02 15:04:34"))
	time.Sleep(2 * time.Second)
}

func MyOperator1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("MyOperator1 Done")
			return
		default:
			PrintUtil("MyOperator1")
		}
	}
}

func MyOperator2(ctx context.Context) {
	fmt.Println("MyOperator2")
}

func MyDo2(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go MyOperator1(ctx, wg)
	go MyOperator2(ctx)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("MyDo2 Done")
			return
		default:
			PrintUtil("MyDo2")
		}
	}
}

func MyDo1(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go MyDo2(ctx, wg)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("MyDo1 Done")
			// 打印 ctx 关闭原因
			fmt.Println(ctx.Err())
			return
		default:
			PrintUtil("MyDo1")
		}
	}
}

func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go MyDo1(ctx, &wg)

	time.Sleep(5 * time.Second)
	fmt.Println("stop all goroutines")
	cancel()

	wg.Wait()
	fmt.Println("all goroutines canceled")
}

func dl2(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("dl2 : ", n)
			n++
			time.Sleep(time.Second)
		}
	}
}

func dl1(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("dl1 : ", n)
			n++
			time.Sleep(2 * time.Second)
		}
	}
}

func Deadline() {
	// 设置deadline为当前时间之后的5秒那个时刻
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	go dl1(ctx)
	go dl2(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("over", ctx.Err())
			return
		}
	}
}

func to1(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("to1 is over")
			return
		default:
			fmt.Println("to1 : ", n)
			n++
			time.Sleep(time.Second)
		}
	}
}
func WithTimeout() {
	// 设置为6秒后context结束
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	go to1(ctx)
	n := 1
	for {
		select {
		case <-time.Tick(2 * time.Second):
			if n == 9 {
				return
			}
			fmt.Println("number :", n)
			n++
		}
	}
}

func v3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("v3 Done : ", ctx.Err())
			return
		default:
			fmt.Println(ctx.Value("key"))
			time.Sleep(3 * time.Second)
		}
	}
}
func v2(ctx context.Context) {
	fmt.Println(ctx.Value("key"))
	fmt.Println(ctx.Value("v1"))
	// 相同键,值覆盖
	ctx = context.WithValue(ctx, "key", "modify from v2")
	go v3(ctx)
}
func v1(ctx context.Context) {
	if v := ctx.Value("key"); v != nil {
		fmt.Println("key = ", v)
	}
	ctx = context.WithValue(ctx, "v1", "value of v1 func")
	go v2(ctx)
	for {
		select {
		default:
			fmt.Println("print v1")
			time.Sleep(time.Second * 2)
		case <-ctx.Done():
			fmt.Println("v1 Done : ", ctx.Err())
			return
		}
	}
}
func WithValue() {
	ctx, cancel := context.WithCancel(context.Background())
	// 向context中传递值
	ctx = context.WithValue(ctx, "key", "main")
	go v1(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

func main() {
	WithValue()
}
