package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrorNotFound         = status.Errorf(codes.NotFound, "Contact not found")
	ErrorDuplicateContact = status.Errorf(codes.AlreadyExists, "Contact already exists")
)
