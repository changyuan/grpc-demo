
## grpc demo

```bash
cd pbfiles
protoc --go_out=plugins=grpc:../services Prod.proto


# run server
go run server/main.go

# run client
go run client/main.go
```