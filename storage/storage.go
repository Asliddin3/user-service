package storage

import (
	"github.com/Asliddin3/user-servis/storage/postgres"
	"github.com/Asliddin3/user-servis/storage/repo"
	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	User() repo.UserStorageI
}

type storagePg struct {
	db           *sqlx.DB
	customerRepo repo.UserStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:           db,
		customerRepo: postgres.NewUserRepo(db),
	}
}
func (s storagePg) User() repo.UserStorageI {
	return s.customerRepo
}
