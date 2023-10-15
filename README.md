# About

this is a golang exercise using:
* Golang
* gRPC
* mongoDB


```
protoc -Iblog/proto --go_opt=module=github.com/rafaelmbcosta/golang-blog-experiment --go_out=. --go-grpc_opt=module=github.com/rafaelmbcosta/golang-blog-experiment --go-grpc_out=. blog/proto/*.proto
```