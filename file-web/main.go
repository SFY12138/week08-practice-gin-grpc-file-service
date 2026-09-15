package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	pb "file-web/proto/file"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()

	// 创建 uploads 目录
	os.MkdirAll("./uploads", os.ModePerm)

	r.POST("/upload", uploadHandler)
	r.GET("/files", getFilesHandler)

	fmt.Println("Gin Web Server is running on port :8080")
	r.Run(":8080")
}

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// 打开文件计算 MD5
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate hash"})
		return
	}
	fileHash := fmt.Sprintf("%x", hash.Sum(nil))

	// 重新打开文件保存到本地
	src, err = file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reopen file"})
		return
	}
	defer src.Close()

	// 保存文件
	dstPath := filepath.Join("./uploads", file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// 调用 gRPC 文件服务
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to gRPC server: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to file service"})
		return
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)
	uploadTime := time.Now().Format(time.RFC3339)

	resp, err := client.SaveFile(c, &pb.FileRequest{
		Filename:   file.Filename,
		Hash:       fileHash,
		Size:       file.Size,
		UploadTime: uploadTime,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":    resp.Success,
		"message":    resp.Message,
		"file_id":    resp.FileId,
		"filename":   file.Filename,
		"hash":       fileHash,
		"size":       file.Size,
		"uploadTime": uploadTime,
	})
}

func getFilesHandler(c *gin.Context) {
	page := 1
	pageSize := 10

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if ps := c.Query("page_size"); ps != "" {
		fmt.Sscanf(ps, "%d", &pageSize)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to file service"})
		return
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)
	resp, err := client.GetFileList(c, &pb.FileListRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var files []gin.H
	for _, f := range resp.Files {
		files = append(files, gin.H{
			"id":         f.Id,
			"filename":   f.Filename,
			"hash":       f.Hash,
			"size":       f.Size,
			"uploadTime": f.UploadTime,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"message": resp.Message,
		"total":   resp.Total,
		"page":    page,
		"pageSize": pageSize,
		"files":   files,
	})
}
