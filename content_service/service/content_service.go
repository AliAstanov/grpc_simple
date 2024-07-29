package service

import (
	"content_service/genproto/content_service"
	storage "content_service/storage"
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
)

type ContentService struct {
	Storage storage.StorageI
	content_service.UnimplementedContentServiceServer
}

func NewContentService(storage storage.StorageI) *ContentService {
	return &ContentService{Storage: storage}
}

func (c *ContentService) CreateContent(ctx context.Context, req *content_service.CreateContentReq) (*content_service.Content, error) {
	c.Storage.GetContentRepo().Create(ctx, req)

	return &content_service.Content{
		Id:        uuid.New().String(),
		Title:     req.Title,
		CreatedAt: time.Now().GoString(),
	}, nil
}

func (c *ContentService) GetContentList(ctx context.Context, req *content_service.GetListReq) (*content_service.GetListResp, error) {
	return c.Storage.GetContentRepo().GetList(ctx, req)
}

func (c *ContentService) GetContentById(ctx context.Context, req *content_service.GetByIdReq) (*content_service.Content, error) {
	return c.Storage.GetContentRepo().GetContentById(ctx, req.Id)
}
func (c *ContentService) UpdateContent(ctx context.Context, req *content_service.UpdateContentReq) (*empty.Empty, error) {
	log.Println("er---888", req)

	return c.Storage.GetContentRepo().Update(ctx, req)
}
func (c *ContentService) DeleteContent(ctx context.Context, id *content_service.DeleteContentReq) (*empty.Empty, error) {
	return c.Storage.GetContentRepo().Delete(ctx, id)
}
