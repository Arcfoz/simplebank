package gapi

import (
	db "github.com/arcfoz/simplebank/db/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/arcfoz/simplebank/pb"
)

func convertUser(user db.Users) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
