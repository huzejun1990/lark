protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  *.proto

#protoc --go_out=. --go_opt=paths=source_relative \
#  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
#  *.proto


#protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_out=paths=source_relative *.proto


#protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative $proto