# About

this is a golang exercise using:
* Golang
* gRPC
* mongoDB


# setting up
```
protoc -Iblog/proto --go_opt=module=github.com/rafaelmbcosta/golang-blog-experiment --go_out=. --go-grpc_opt=module=github.com/rafaelmbcosta/golang-blog-experiment --go-grpc_out=. blog/proto/*.proto
go build -o bin/blog/server ./blog/server
go build -o bin/blog/client ./blog/client
```

# running
```
docker-compose up
./bin/blog/server
./bin/blog/client
```