package grpc_phonebook

//go:generate protoc -I=. --js_out=import_style=commonjs,binary:. --go_out=plugins=grpc,import_path=$GOPATH/src/github.com/iheanyi/grpc-phonebook:. phonebook.proto
