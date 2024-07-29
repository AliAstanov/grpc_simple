package service

import (
	"api_gateway/genproto/content_service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Service()content_service.ContentServiceClient{
	conn,err := grpc.NewClient("localhost:8000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("did not connect:",err)
		return nil
	}

	contentService := content_service.NewContentServiceClient(conn)

	return contentService
}