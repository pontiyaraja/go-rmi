package client

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "grpc_blog/model/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	fmt.Println("client connected to server.")
	client := pb.NewHelloServiceClient(conn)
	ctx := context.Background()
	blog, err := client.Create(ctx, &pb.Blog{PostID: 1, Content: "Test blog", Author: "I'm", Tags: "test", PublicationDate: "my Publication"})
	fmt.Println("Create Response : ", blog, "error", err)
	blog, err = client.Read(ctx, &pb.PostID{PostID: 1})
	fmt.Println("Read Response : ", blog, "error", err)
	blog, err = client.UPdate(ctx, &pb.Blog{PostID: 1, Content: "Test blog Updated", Author: "I'm", Tags: "test", PublicationDate: "my Publication"})
	fmt.Println("Update Response : ", blog, "error", err)
	blog, err = client.Read(ctx, &pb.PostID{PostID: 1})
	fmt.Println("Read Response : ", blog, "error", err)
	blog, err = client.Delete(ctx, &pb.PostID{PostID: 1})
	fmt.Println("Delete Response : ", blog, "error", err)
	blog, err = client.Read(ctx, &pb.PostID{PostID: 1})
	fmt.Println("Read Response : ", blog, "error", err)
	fmt.Println("ending session with server ...........")
}
