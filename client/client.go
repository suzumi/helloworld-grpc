package main
 
import (
    "context"
    "fmt"
    "log"
 
    pb "github.com/suzumi/helloworld-grpc/pb"
 
    "google.golang.org/grpc"
)
 
func main() {
    //sampleなのでwithInsecure
    conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
    if err != nil {
        log.Fatal("client connection error:", err)
    }
    defer conn.Close()
    client := pb.NewGreeterClient(conn)
    message := &pb.HelloRequest{"suzumi"}
    res, err := client.SayHello(context.TODO(), message)
    fmt.Printf("result:%#v \n", res)
    fmt.Printf("error::%#v \n", err)
}