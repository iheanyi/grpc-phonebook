package api

//go:generate protoc -I=. --proto_path=${GOPATH}/src:. --go_out=plugins=grpc:. --js_out=import_style=commonjs,binary:. api.proto
