
## Pre-requisites
Install protobuff compiler
```bash
brew install protobuf
protoc --version
```

Install protoc-gen-go
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Install protoc-gen-go-grpc
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

```


```bash
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.bash_profile
```