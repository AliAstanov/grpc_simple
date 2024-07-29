package postgres

import (
	"content_service/genproto/content_service"
	halpers "content_service/pkg/halper"
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ContentRepo struct {
	conn *pgx.Conn
}

func NewContentRepo(conn *pgx.Conn) ContentRepoI {
	return &ContentRepo{conn: conn}
}

func (c *ContentRepo) Create(ctx context.Context, contentreq *content_service.CreateContentReq) (*content_service.Content, error) {
	var err error
	var content = &content_service.Content{}
	content.Id = uuid.New().String()
	content.CreatedAt = time.Now().Format(time.RFC3339)

	err = halpers.DataParser1(contentreq, content)
	if err != nil {
		return nil, err
	}

	url := `
		INSERT INTO 
			content(
				id,
				title,
				created_at
			)VALUES(
				$1,$2,$3
			)`

	_, err = c.conn.Exec(
		ctx, url,
		content.Id,
		content.Title,
		content.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	log.Println("content:", content)
	return content, nil
}

func (c *ContentRepo) GetList(ctx context.Context, req *content_service.GetListReq) (*content_service.GetListResp, error) {
	var contents = &content_service.GetListResp{
		Contents: []*content_service.Content{},
	}
	query := `
		SELECT 
			*
		FROM
			content
		LIMIT
			 $1
		OFFSET
			$2
	`
	offset := (req.Offset - 1) * req.Limit
	rows, err := c.conn.Query(
		ctx, query,
		req.Limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var content content_service.Content
		var createdAt time.Time

		if err := rows.Scan(&content.Id, &content.Title, &createdAt); err != nil {
			log.Println("error on rows.Scan:", err)
			return nil, err
		}

		content.CreatedAt = createdAt.Format(time.RFC3339)
		contents.Contents = append(contents.Contents, &content)
	}

	var count int32
	err = c.conn.QueryRow(ctx, "SELECT count(*) FROM content").Scan(&count)
	if err != nil {
		log.Fatal("error on scanning content count.")
		return nil, err
	}
	defer rows.Close()

	return &content_service.GetListResp{
		Contents: contents.Contents,
		Count:    count,
	}, nil
}
func (c *ContentRepo) GetContentById(ctx context.Context, id string) (*content_service.Content, error) {

	var content content_service.Content

	query := `
		SELECT 
			*
		FROM
			content
		WHERE
			id = $1
	`
	row := c.conn.QueryRow(ctx, query, id)

	if err := row.Scan(
		&content.Id,
		&content.Title,
		&content.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("content not found")
		}
	}

	return &content, nil
}
func (c *ContentRepo) Update(ctx context.Context, content *content_service.UpdateContentReq) (*empty.Empty, error) {
	log.Println("Starting update operation")

	var updatedContent content_service.Content
	query := `
		UPDATE 
			content
		SET 
			title = $1
		WHERE 
			id = $2
	`

	_, err := c.conn.Exec(
		ctx, query,
		content.Title,
		content.Id,
	)
	if err != nil {
		log.Println("No rows found for the given ID", err)
		return nil, err
	}

	log.Printf("Update operation successful: %v", &updatedContent)
	return &emptypb.Empty{}, nil
}

func (c *ContentRepo) Delete(ctx context.Context, req *content_service.DeleteContentReq) (*empty.Empty, error) {
	log.Println("sss")
	query := `
		DELETE FROM
			content
		WHERE
			id = $1
	`
	_, err := c.conn.Exec(ctx, query, req.Id)
	log.Println("ssrrs")
	if err != nil {
		log.Println("error on Delete Content:", err)
		return nil, err
	}
	log.Println("ss22s")
	return &emptypb.Empty{}, nil
}
