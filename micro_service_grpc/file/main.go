package main

import (
	"context"
	"file/server/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedArchiveServer
}

func (s *server) GetFileList(ctx context.Context, in *pb.GetFileListRequest) (*pb.GetFileListResponse, error) {
	return &pb.GetFileListResponse{
		Files: getFileList(),
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

func getFileList() []*pb.File {
	fileList := []*pb.File{{FileName: "file_2.extension"},
		{FileName: "file_1.extension"},
		{FileName: "file_3.extension"}}
	return fileList
}
