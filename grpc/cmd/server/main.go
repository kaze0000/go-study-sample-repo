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
	// hello_grpc.pb.goに以下のようにある
	// UnimplementedGreetingServiceServer must be embedded to have forward compatible implementations.
	hellopb.UnimplementedGreetingServiceServer
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()), //GetNameもhello_grpc.pb.goに書かれている
	}, nil
}

//  自作サービス構造体のコンストラクタを定義(コンストラクタ関数の名前は、Newで始まる)
func NewMyServer() *myServer {
	return &myServer{} //{}は中身が空のばあいフィールドがゼロ値になる。つまり、初期化している
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
	// ↓hellopb.RegisterGreetingServiceServer(s, [サーバーに登録するサービス])
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())
	// ↑の第2引数には、GreetingServiceServer interface(hello_grpc.pb.goにある)を実装した構造体が入る
	// ので、第二引数に代入できる自作構造体を定義し、そこでHelloメソッドを実装する。それが上記のmyServer

	reflection.Register(s) //gRPCurlを使うためには、リクエストを送るgRPCサーバーに「サーバーリフレクション」という設定がなされていることが前提となるので、ここで設定
												 //gRPCulrは、シリアライズのルール知らないので、gRPCサーバから取得できるようにする設定

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
