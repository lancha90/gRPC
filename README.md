# Serializing Methods - gRPC


Hello gRPC


## Generate Struts && Service Interface

protoc -I=$SRC_DIR/proto --go-grpc_out=$SRC_DIR/person --go_out=$SRC_DIR/person $SRC_DIR/proto/person.proto

protoc -I=$SRC_DIR/proto --go-grpc_out=$SRC_DIR --go_out=$SRC_DIR $SRC_DIR/proto/*.proto