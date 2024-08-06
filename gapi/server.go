package gapi

import (
	"fmt"

	db "github.com/arcfoz/simplebank/db/sqlc"
	"github.com/arcfoz/simplebank/pb"
	"github.com/arcfoz/simplebank/token"
	"github.com/arcfoz/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
