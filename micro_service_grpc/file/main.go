package main

import (
	"context"
	"file/server/pb"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedArchiveServer
}

func (s *server) GetBookList(ctx context.Context, in *pb.GetFileListRequest) (*pb.GetFileListResponse, error) {
	return &pb.GetFileListResponse{
		File: getFileList(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8100")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterArchiveServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
