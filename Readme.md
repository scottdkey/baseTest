install `protoc` complier

```bash
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

Generate Code with

```bash
protoc -I proto/ proto/service.proto --go_out=plugins=grpc:.
```