# gRPC Simple Service Tutorial


## Generate Struts && Service Interface

```
protoc -I=$SRC_DIR/proto --go-grpc_out=$SRC_DIR --go_out=$SRC_DIR $SRC_DIR/proto/*.proto
```

## Run Service

```
go run main.go
```