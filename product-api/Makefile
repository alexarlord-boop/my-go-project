check_install:
	which swagger || GO111MODULE=on go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger: check_install
	GO111MODULE=on swagger generate spec -o ./swagger.yaml --scan-models