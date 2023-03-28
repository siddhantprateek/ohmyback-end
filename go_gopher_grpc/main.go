package main

import "context"

type server struct {
	pb.UnimplementedInventoryServer
}

func (s *server) GetBooks(ctx context.Context, in *pb.GetBooks) {

}
