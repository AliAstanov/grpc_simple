package service

import (
	"content_service/config"
	"content_service/genproto/content_service"
	db "content_service/pkg/postgres"
	storage "content_service/storage"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Service() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
		return
	}

	server := grpc.NewServer()

	cfg := config.Load()

	conn, err := db.ConnectToDb(cfg.PgConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	storage := storage.NewStorage(conn)

	contentService := NewContentService(storage)

	content_service.RegisterContentServiceServer(server, contentService)

	log.Println("gRPC server is running on port 8000")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
