package storage

import (
	"test-ozon/internal/storage/config"
	"test-ozon/internal/storage/memory"
	"test-ozon/internal/storage/postgres"
)


type DB struct {
    DbInterface
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=dbInterface
type DbInterface interface {
	GetUrl(aliasUrl string) (string, error)
    SaveUrl(urlToSave, aliasUrl string) error
    IsDublicate(alias string) bool
}

func NewDB(name string) (*DB, error) {
    var (
		db DbInterface
		err error
	)

    switch name {
    case "memory":
		db = memory.NewMemory()
    case "postgres":
		cfg := config.NewConfig_Postgres()
        db, err = postgres.NewPostgres(cfg)
    default:
        cfg := config.NewConfig_Postgres()
        db, err = postgres.NewPostgres(cfg)
    }

    return &DB{
        DbInterface: db,
    }, err
}