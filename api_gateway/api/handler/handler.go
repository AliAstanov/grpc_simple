package handler

import "api_gateway/genproto/content_service"

type handler struct {
	conn content_service.ContentServiceClient
}

func NewHandler(conn content_service.ContentServiceClient) *handler {
	return &handler{conn: conn}
}
