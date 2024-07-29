package handler

import (
	"api_gateway/genproto/content_service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateContent(ctx *gin.Context) {
	var reqBody content_service.CreateContentReq

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println("error binding request body:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.conn.CreateContent(ctx, &reqBody)
	if err != nil {
		log.Println("error creating content:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create content"})
		return
	}

	ctx.JSON(201, resp)
}
func (h *handler) GetContents(ctx *gin.Context) {
	var reqBody content_service.GetListReq
	limit := ctx.Query("limit")
	page := ctx.Query("page")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("Invalid limit parameter:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}
	reqBody.Limit = int32(limitInt)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Invalid page parameter:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}
	reqBody.Offset = int32(pageInt)

	contents, err := h.conn.GetContentList(ctx, &reqBody)
	if err != nil {
		log.Println("Failed to fetch content list:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch content list"})
		return
	}
	ctx.JSON(http.StatusOK, contents)
}

func (h *handler) GetContentById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		log.Println("ID is required")
		ctx.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	content, err := h.conn.GetContentById(ctx, &content_service.GetByIdReq{Id: id})
	if err != nil {
		log.Println("error fetching content by ID:", err)
		ctx.JSON(500, gin.H{"error": "Failed to fetch content"})
		return
	}

	ctx.JSON(200, content)
}

func (h *handler) UpdateContent(ctx *gin.Context) {
	var reqBody content_service.UpdateContentReq
	id := ctx.Param("id")
	if id == "" {
		log.Println("ID is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		log.Println("Invalid request body:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := &content_service.UpdateContentReq{
		Id:    id,
		Title: reqBody.Title,
	}

	updatedContent, err := h.conn.UpdateContent(ctx, req)
	if err != nil {
		log.Println("Error updating content:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update content"})
		return
	}

	ctx.JSON(http.StatusOK, updatedContent)
}
func (h *handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		log.Println("ID is required")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	req := &content_service.DeleteContentReq{
		Id: id,
	}

	_, err := h.conn.DeleteContent(ctx, req)
	if err != nil {
		log.Println("Error deleting content:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete content"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Content deleted successfully"})
}
