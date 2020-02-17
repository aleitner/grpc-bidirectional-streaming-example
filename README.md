# grpc-bidirectional-stream

Provide Example `client` and `server` code for bidirectional streaming using grpc.

## Dependencies
* protoc-gen-go
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```


## Build 

### Protobuf

```bash
cd pkg/proto
protoc *.proto --go_out=plugins=grpc:.
```

## Run

### Server
* Run the server process
```bash
go run cmd/server/main.go
```

### Client
* Interact with the server process
```bash
go run cmd/client/main.go
```