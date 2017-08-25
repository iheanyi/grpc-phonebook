package api

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	ErrorNotFound         = grpc.Errorf(status.NotFound, "Contact not found")
	ErrorDuplicateContact = grpc.Errorf(status.AlreadyExists, "Contact already exists")
)
