package main

import (
	"time"
	"fmt"
	"context"
	"configcenter/src/framework/core/errors"
)

var key string = "name"

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	////附加值
	//valueCtx := context.WithValue(ctx, key, "【监控1】")
	//go watch(valueCtx)
	//time.Sleep(10 * time.Second)
	//fmt.Println("可以了，通知监控停止")
	//cancel()
	////为了检测监控过是否停止，如果没有监控输出，就表示停止了
	//time.Sleep(5 * time.Second)

	fmt.Println("my test")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 3*time.Second)
	ctx2 := context.WithValue(ctx1, "miya", "test")
	go doSomething(ctx2)
	time.Sleep(4 * time.Second)
	cancel1()
	fmt.Println("my test over")

	ctx3, cancel3 := context.WithCancel(context.Background())
	v, err := slowOperationWithTimeout(ctx3)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(v)
	}
	time.Sleep(2000 * time.Millisecond)
	cancel3()

}

func slowOperationWithTimeout(ctx context.Context) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel() // releases resources if slowOperation completes before timeout elapses
	return slowOperation(ctx)
}

func slowOperation(ctx context.Context) (interface{}, error) {
	fmt.Println("start", ctx.Value("nothing"))
	for {

		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value("miya"), "timeout")
			return nil, errors.New("timeout")
		default:
			fmt.Println(ctx.Value("miya"), "running...")
			return 1, nil
			time.Sleep(2000 * time.Millisecond)
		}

	}
	fmt.Println("over")

	return nil, nil
}

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value("miya"), "timeout")
			return
		default:
			fmt.Println(ctx.Value("miya"), "default...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
