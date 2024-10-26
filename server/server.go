package main

import (
	"context"
	
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	

	pb "github.com/RonexLemon/Transfer/service/genproto"
)

type Server struct{
	pb.UnimplementedTransferServiceServer

}
type txData struct{
	Name string
	Amount  string
	From string
}

func hashTxData(data txData) (string, error) {
	
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	
	hash := sha256.Sum256(jsonData)

	
	return hex.EncodeToString(hash[:]), nil
}

func (s*Server) Transfer(ctx context.Context,request *pb.TransferRequest) (*pb.TransferResponse, error) {
	txdata   := txData{Name: request.Amount,Amount: request.Amount,From: request.From}

	txhash,err := hashTxData(txdata)
	if err !=nil{
		return nil,err
	}


	
	
	return &pb.TransferResponse{Hash: txhash}, nil
}