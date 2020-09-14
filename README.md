
## grpc demo

```bash
cd pbfiles
protoc --go_out=plugins=grpc:../services Product.proto

# 或使用 pb 文件增加
option go_package = "../services";
# 执行
protoc --go_out=plugins=grpc:./services pbfiles/Product.proto


# run server
go run server/main.go

# run client
go run client/main.go
```