package file

import (
	"context"
	"google.golang.org/grpc"
)

type FileServiceClient interface {
	SaveFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error)
	GetFileList(ctx context.Context, in *FileListRequest, opts ...grpc.CallOption) (*FileListResponse, error)
}

type fileServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileServiceClient(cc *grpc.ClientConn) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) SaveFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/SaveFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFileList(ctx context.Context, in *FileListRequest, opts ...grpc.CallOption) (*FileListResponse, error) {
	out := new(FileListResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/GetFileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
