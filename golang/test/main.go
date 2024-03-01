package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work canceled")
			return
		default:
			// Do some work
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func doCleanup(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Cleanup canceled")
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go doWork(ctx)
	go doCleanup(ctx)

	// 模拟运行一段时间后取消
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
