package config

import (
	"tfg/blockchain-service/pkg/domain/user"
	user2 "tfg/blockchain-service/pkg/infraestructure/user"
	"tfg/blockchain-service/pkg/libs/sql"
)

type Dependencies struct {
	UserRepository user.Repository
}

func BuildDependencies(config *Config) (*Dependencies, error) {
	db, err := sql.NewDB(config.DBConfig)

	if err != nil {
		return nil, err
	}
	userRepository := user2.NewPostgresUserRepository(db)
	return &Dependencies{
		userRepository,
	}, nil
}
