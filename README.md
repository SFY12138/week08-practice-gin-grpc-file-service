# 文件上传系统

基于 Gin + gRPC + SQLite 的文件上传系统，包含 Web 服务和文件服务两个独立模块。

## 技术栈

- **Web 服务**: Go + Gin
- **文件服务**: Go + gRPC + SQLite
- **协议定义**: Protocol Buffers

## 项目结构

```
gin-grpc-file-service/
├── file-web/          # Gin Web 服务
│   ├── main.go        # 入口文件
│   ├── go.mod         # Go 依赖管理
│   ├── go.sum         # 依赖锁文件
│   ├── uploads/       # 文件上传目录
│   └── proto/         # gRPC 协议定义
├── file-service/      # gRPC 文件服务
│   ├── main.go        # 入口文件
│   ├── go.mod         # Go 依赖管理
│   ├── go.sum         # 依赖锁文件
│   ├── files.db       # SQLite 数据库
│   └── proto/         # gRPC 协议定义
├── proto/             # 原始协议定义
│   └── file.proto     # gRPC 协议文件
└── README.md          # 项目说明文档
```

## 功能说明

### Web 服务 (file-web)

负责接收 HTTP 请求，处理文件上传：
1. 接收用户上传的文件
2. 保存文件到 `uploads` 目录
3. 计算文件 MD5 hash
4. 调用 gRPC 文件服务保存文件记录

### 文件服务 (file-service)

提供 gRPC 接口，管理文件记录：
1. `SaveFile`: 保存文件记录到 SQLite
2. `GetFileList`: 查询文件列表（支持分页）

## 运行方式

### 1. 启动 gRPC 文件服务

```bash
cd file-service
go run main.go
```

服务运行在 `localhost:50051`

### 2. 启动 Web 服务

```bash
cd file-web
go run main.go
```

服务运行在 `localhost:8080`

## API 接口

### 上传文件

```bash
POST /upload
Content-Type: multipart/form-data

curl -X POST -F "file=@test.txt" http://localhost:8080/upload
```

### 查询文件列表

```bash
GET /files?page=1&page_size=10

curl http://localhost:8080/files?page=1&page_size=10
```

## gRPC 接口定义

```protobuf
service FileService {
  rpc SaveFile(FileRequest) returns (FileResponse);
  rpc GetFileList(FileListRequest) returns (FileListResponse);
}
```

## 数据流程

```
用户上传文件 → HTTP → Gin → 保存文件 → 计算Hash → gRPC → SQLite
                                         ↓
                                   返回上传结果
```

## 注意事项

1. 确保 gRPC 文件服务先启动
2. `uploads` 目录会自动创建
3. SQLite 数据库文件会自动创建
4. 文件 hash 用于去重，相同 hash 的文件不会重复保存