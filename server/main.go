package main
 
import (
    "log"
    "net"
 
    pb "github.com/suzumi/helloworld-grpc/pb"
    "github.com/suzumi/helloworld-grpc/server/service"
 
    "google.golang.org/grpc"
)
 
func main() {
    listenPort, err := net.Listen("tcp", ":19003")
    if err != nil {
        log.Fatalln(err)
    }
    server := grpc.NewServer()
    helloService := &service.MyHelloService{}
    // 実行したい実処理をseverに登録する
    pb.RegisterGreeterServer(server, helloService)
    server.Serve(listenPort)
}