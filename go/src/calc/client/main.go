package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "app/dict"
)

func main() {

	conn, err := grpc.Dial("localhost:9001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("error to connect: %+v", err)
	}

	c := pb.NewTranslateServiceClient(conn)

	add, _ := c.CalcAdd(context.Background(), &pb.AddRequest{Num1: 12, Num2: 4})
	fmt.Println("+",add)

	div, _ := c.CalcDiv(context.Background(), &pb.DivRequest{Num1: 12, Num2: 4})
	fmt.Println("-",div)

	sub, _ := c.CalcSub(context.Background(), &pb.SubRequest{Num1: 12, Num2: 4})
	fmt.Println("/",sub)

	mult, err := c.CalcMult(context.Background(), &pb.MultRequest{Num1: 12, Num2: 4})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("*",mult)

	min, err := c.CalcMin(context.Background(), &pb.MinRequest{Nums: []int32{4,6,8,3,14}})

	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("min",min.Res)

	sqrt, err := c.CalcSqrt(context.Background(), &pb.SqrtRequest{Num: 9})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sqrt",sqrt)
	
	
	pow, err := c.CalcPow(context.Background(), &pb.PowRequest{Num1: 2, Num2: 7})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pow",pow)

}
