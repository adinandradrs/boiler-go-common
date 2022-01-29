package base

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type BaseRepositoryCapsule struct {
	Database *pgxpool.Pool
}

func BaseRepository(database *pgxpool.Pool) *BaseRepositoryCapsule {
	return &BaseRepositoryCapsule{
		Database: database,
	}
}

type BaseRepositoryInterface interface {
	Save(input interface{}) (interface{}, error)
	FindById(id uint) (interface{}, error)
	Delete(id uint) error
}

type BaseServiceInterface interface {
	Execute(input interface{}) (interface{}, error)
}
