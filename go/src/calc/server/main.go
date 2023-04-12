package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	pb "app/dict"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTranslateServiceServer
}


func (s *server) CalcAdd(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{Res: req.Num1 + req.Num2},nil
}

func (s *server) CalcDiv(ctx context.Context, req *pb.DivRequest) (*pb.DivResponse, error) {
	return &pb.DivResponse{Res: req.Num1 - req.Num2},nil
}

func (s *server) CalcSub(ctx context.Context, req *pb.SubRequest) (*pb.SubResponse, error) {
	return &pb.SubResponse{Res: req.Num1 / req.Num2},nil
}

func (s *server) CalcMult(ctx context.Context, req *pb.MultRequest) (*pb.MultResponse, error) {
	return &pb.MultResponse{Res: req.Num1 * req.Num2},nil
}
func (s *server) CalcMin(ctx context.Context, req *pb.MinRequest) (*pb.MinResponse, error) {

var	min int32 
min = req.Nums[0]
	for _, v := range req.Nums{
		if  v< min {
			min = v
		}
		if min <=0 {
			return &pb.MinResponse{},nil
		}
	}
	return &pb.MinResponse{Res: min},nil
}

func (s *server) CalcSqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	
	return &pb.SqrtResponse{Res: int32(math.Sqrt(float64(req.Num)))},nil
}
func (s *server) CalcPow(ctx context.Context, req *pb.PowRequest) (*pb.PowResponse, error) {

	return &pb.PowResponse{Res: int32(math.Pow(float64(req.Num1), float64(req.Num2)))},nil
}






func main() {

	lis, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTranslateServiceServer(s, &server{})

	fmt.Println("Listen RPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %+v", err)
	}
}
