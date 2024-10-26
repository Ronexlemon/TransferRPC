package main

import (
	"context"
	"log"
	
	"time"

	pb "github.com/RonexLemon/Transfer/service/genproto"
	"google.golang.org/grpc"
)

func main(){
	// Create a new client
	//client := pb.NewTransferServiceClient("localhost:9090")
	
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Failed to listen")
	}

	defer conn.Close()
	client:= pb.NewTransferServiceClient(conn)
	data:=  &pb.TransferRequest{From: "0x3455657376377683768345787565723657365736756376",To: "0x362763756375673562375637256372562375672367",Amount: "100",BlocktimeStamp:time.Second.String()}
	ctx,cancel:= context.WithTimeout(context.Background(),time.Second);
	defer cancel()
	txHash,err := client.Transfer(ctx,data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(txHash)
	


}