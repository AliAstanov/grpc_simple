package storage

import (
	"content_service/storage/postgres"

	"github.com/jackc/pgx/v5"
)

type StorageI interface {
	GetContentRepo() postgres.ContentRepoI
}

type storage struct {
	contentRepo postgres.ContentRepoI
}

func NewStorage(conn *pgx.Conn) StorageI {
	return &storage{
		contentRepo: postgres.NewContentRepo(conn),
	}
}

func (s *storage) GetContentRepo() postgres.ContentRepoI {
	return s.contentRepo
}
