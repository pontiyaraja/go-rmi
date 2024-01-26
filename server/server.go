/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC server that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It implements the route guide service whose definition can be found in routeguide/route_guide.proto.
package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "grpc_blog/model/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type blogServer struct {
	pb.UnimplementedHelloServiceServer
	blogs map[int32]*pb.Blog
}

func newServer() *blogServer {
	s := &blogServer{blogs: make(map[int32]*pb.Blog)}
	return s
}

func main() {
	fmt.Println("Starting server .........")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func (server *blogServer) Create(ctx context.Context, blog *pb.Blog) (*pb.Blog, error) {
	if blog == nil {
		return nil, status.Errorf(codes.InvalidArgument, "blog is nil")
	}
	if _, ok := server.blogs[blog.PostID]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "duplicate blog")
	}
	server.blogs[blog.PostID] = blog
	return blog, nil
}

func (server *blogServer) Read(ctx context.Context, postId *pb.PostID) (*pb.Blog, error) {
	if postId == nil {
		return nil, status.Errorf(codes.InvalidArgument, "post ID is nil")
	}
	blog, ok := server.blogs[postId.PostID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "blog not found")
	}
	return blog, nil
}

func (server *blogServer) UPdate(ctx context.Context, blog *pb.Blog) (*pb.Blog, error) {
	if blog == nil {
		return nil, status.Errorf(codes.InvalidArgument, "blog is nil")
	}
	if _, ok := server.blogs[blog.PostID]; !ok {
		return nil, status.Errorf(codes.NotFound, "no blog found to update")
	}
	server.blogs[blog.PostID] = blog
	return blog, nil
}

func (server *blogServer) Delete(ctx context.Context, postId *pb.PostID) (*pb.Blog, error) {
	if postId == nil {
		return nil, status.Errorf(codes.InvalidArgument, "post ID is nil")
	}
	blog, ok := server.blogs[postId.PostID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "no blog found to delete")
	}
	delete(server.blogs, blog.PostID)
	return blog, nil
}
