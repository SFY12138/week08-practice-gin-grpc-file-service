package file

import (
	"context"

	"google.golang.org/grpc"
)

type FileServiceServer interface {
	SaveFile(ctx context.Context, req *FileRequest) (*FileResponse, error)
	GetFileList(ctx context.Context, req *FileListRequest) (*FileListResponse, error)
}

type UnimplementedFileServiceServer struct{}

func (UnimplementedFileServiceServer) SaveFile(ctx context.Context, req *FileRequest) (*FileResponse, error) {
	return nil, nil
}

func (UnimplementedFileServiceServer) GetFileList(ctx context.Context, req *FileListRequest) (*FileListResponse, error) {
	return nil, nil
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveFile",
			Handler:    _FileService_SaveFile_Handler,
		},
		{
			MethodName: "GetFileList",
			Handler:    _FileService_GetFileList_Handler,
		},
	},
}

func _FileService_SaveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).SaveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/SaveFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).SaveFile(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/GetFileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFileList(ctx, req.(*FileListRequest))
	}
	return interceptor(ctx, in, info, handler)
}
