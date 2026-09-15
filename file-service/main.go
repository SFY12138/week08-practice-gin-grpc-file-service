package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "file-service/proto/file"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
	db *sql.DB
}

func (s *server) SaveFile(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
	result, err := s.db.Exec(
		"INSERT INTO files (filename, hash, size, upload_time) VALUES (?, ?, ?, ?)",
		req.Filename, req.Hash, req.Size, req.UploadTime,
	)
	if err != nil {
		return &pb.FileResponse{Success: false, Message: err.Error()}, nil
	}

	id, _ := result.LastInsertId()
	return &pb.FileResponse{
		Success: true,
		Message: "File saved successfully",
		FileId:  id,
	}, nil
}

func (s *server) GetFileList(ctx context.Context, req *pb.FileListRequest) (*pb.FileListResponse, error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var total int64
	err := s.db.QueryRow("SELECT COUNT(*) FROM files").Scan(&total)
	if err != nil {
		return &pb.FileListResponse{Success: false, Message: err.Error()}, nil
	}

	rows, err := s.db.Query(
		"SELECT id, filename, hash, size, upload_time FROM files LIMIT ? OFFSET ?",
		pageSize, offset,
	)
	if err != nil {
		return &pb.FileListResponse{Success: false, Message: err.Error()}, nil
	}
	defer rows.Close()

	var files []*pb.FileRecord
	for rows.Next() {
		var id int64
		var filename, hash, uploadTime string
		var size int64
		if err := rows.Scan(&id, &filename, &hash, &size, &uploadTime); err != nil {
			return &pb.FileListResponse{Success: false, Message: err.Error()}, nil
		}
		files = append(files, &pb.FileRecord{
			Id:         id,
			Filename:   filename,
			Hash:       hash,
			Size:       size,
			UploadTime: uploadTime,
		})
	}

	return &pb.FileListResponse{
		Success: true,
		Message: "Success",
		Files:   files,
		Total:   total,
	}, nil
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./files.db")
	if err != nil {
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		filename TEXT NOT NULL,
		hash TEXT NOT NULL UNIQUE,
		size INTEGER NOT NULL,
		upload_time TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_files_hash ON files(hash);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &server{db: db})

	fmt.Println("gRPC File Service is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
