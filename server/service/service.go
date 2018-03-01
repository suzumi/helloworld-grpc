package service

import (
	"fmt"
    "context"
    "errors"
    pb "github.com/suzumi/helloworld-grpc/pb"
)
 
type MyHelloService struct {
}
 
func (s *MyHelloService) SayHello(ctx context.Context, message *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{
        Message: fmt.Sprintf("こんにちは、%sさん！", message.Name),
    }, nil
 
    return nil, errors.New("Not Found YourCat")
}
