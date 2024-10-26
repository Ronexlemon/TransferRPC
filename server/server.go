package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	pb "github.com/RonexLemon/Transfer/service/genproto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTransferServiceServer
}
type txData struct {
	Name   string
	Amount string
	From   string
	BlockTimeStamp string
}

func hashTxData(data txData) (string, error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(jsonData)

	return hex.EncodeToString(hash[:]), nil
}

func (s *Server) Transfer(ctx context.Context, request *pb.TransferRequest) (*pb.TransferResponse, error) {
	txdata := txData{Name: request.Amount, Amount: request.Amount, From: request.From,BlockTimeStamp: request.BlocktimeStamp}

	txhash, err := hashTxData(txdata)
	if err != nil {
		return nil, err
	}

	return &pb.TransferResponse{Hash: txhash}, nil
}

func main() {
	fmt.Println("Starting ...")
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTransferServiceServer(grpcServer, &Server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port ... 9090")
}
