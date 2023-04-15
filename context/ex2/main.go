// https://www.youtube.com/watch?v=h2RdcrMLQAo&ab_channel=TutorialEdge](https://www.youtube.com/watch?v=h2RdcrMLQAo&ab_channel=TutorialEdge
package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)
	for {
		select {
		case <-ctx.Done(): //読み取りを待機
				fmt.Println("timed out")
				return
		default:
			fmt.Println("do something cool")
		}
		time.Sleep(500 * time.Millisecond) //0.5秒待機。つまりここが4回実行されて、タイムアウトになる(0.5*4=2秒)
	}
}

func main() {
	fmt.Println("Hello, World!")
	ctx, cancel := context.WithTimeout(context.Background(),2*time.Second) //2秒後にタイムアウトするコンテキストを生成
	defer cancel() //ctxに関連付けられたコンテキストを手動でキャンセルする(つまり、タイムアウトに達する前にキャンセルできる)
	//deferが必要な理由: タイムアウトに達する前に関数が終了した場合でも、ctxに関連付けられたリソースが適切に処分(開放)されることが保証される
	//これは、メモリリークを防げる(リソースが不必要にシステム上に残り続ける現象)
	ctx = enrichContext(ctx)
	go doSomethingCool(ctx)
	select {
	case <- ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
		fmt.Println(ctx.Err()) //コンテキストがキャンセルされたりタイムアウトしたりした場合のエラー情報を返す。コンテキストがまだキャンセルされていない場合はnilを返す
	}
}

// 実行結果
/*
	Hello, World!
	12345
	do something cool
	do something cool
	do something cool
	do something cool
	oh no, I've exceeded the deadline
	context deadline exceeded
*/
