package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "gihub.com/kaze0000/go-study-youtube/grpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

//  自作サービス構造体のコンストラクタを定義
func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// gRPCサーバーを作成
	s := grpc.NewServer()

	// gRPCサーバーにGreetingServiceを登録
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	reflection.Register(s)

	// 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1) // make(chan 型, バッファサイズ) でチャネルを生成。バッファサイズを指定することで、チャネルに格納できる要素数を指定できる
	signal.Notify(quit, os.Interrupt) //第1引数のチャネル(quit)に、第2引数で指定したシグナル(os.Interrupt)が届いたら通知
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop() //リクエストの処理が完了してからサーバーをシャットダウンするための方法で、クライアントとサーバー間で安定した通信を維持し、プロセスを正常に停止するために必要
}
