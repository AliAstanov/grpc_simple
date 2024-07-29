package postgres

import (
	"content_service/genproto/content_service"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type ContentRepoI interface {
	// Create(ctx context.Context, contentreq *models.ContentRequest) (*models.Content, error)
	Create(ctx context.Context, content *content_service.CreateContentReq) (*content_service.Content, error)
	GetList(ctx context.Context, req *content_service.GetListReq) (*content_service.GetListResp, error)
	GetContentById(ctx context.Context, id string) (*content_service.Content, error)
	Update(ctx context.Context, content *content_service.UpdateContentReq) (*empty.Empty, error)
	Delete(ctx context.Context, id *content_service.DeleteContentReq) (*empty.Empty, error)
}
